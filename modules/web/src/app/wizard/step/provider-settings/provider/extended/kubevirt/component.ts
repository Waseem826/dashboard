// Copyright 2024 The Kubermatic Kubernetes Platform contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import {ChangeDetectorRef, Component, forwardRef, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {FormBuilder, NG_VALIDATORS, NG_VALUE_ACCESSOR} from '@angular/forms';
import {ClusterSpecService} from '@core/services/cluster-spec';
import {PresetsService} from '@core/services/wizard/presets';
import {FilteredComboboxComponent} from '@shared/components/combobox/component';
import {KubeVirtCloudSpec} from '@shared/entity/cluster';
import {KubeVirtVPC} from '@shared/entity/provider/kubevirt';
import {NodeProvider} from '@shared/model/NodeProviderConstants';
import {BaseFormValidator} from '@shared/validators/base-form.validator';
import _ from 'lodash';
import {EMPTY, merge, Observable, onErrorResumeNext} from 'rxjs';
import {catchError, debounceTime, filter, map, switchMap, takeUntil, tap} from 'rxjs/operators';

enum Controls {
  Subnet = 'subnet',
  VPCID = 'vpcID',
}

enum SubnetState {
  Loading = 'Loading...',
  Ready = 'Subnet',
  Empty = 'No Subnets Available',
}

enum VPCState {
  Loading = 'Loading...',
  Ready = 'VPC',
  Empty = 'No VPCs Available',
}

@Component({
  selector: 'km-wizard-kubevirt-provider-extended',
  templateUrl: './template.html',
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => KubeVirtProviderExtendedComponent),
      multi: true,
    },
    {
      provide: NG_VALIDATORS,
      useExisting: forwardRef(() => KubeVirtProviderExtendedComponent),
      multi: true,
    },
  ],
})
export class KubeVirtProviderExtendedComponent extends BaseFormValidator implements OnInit, OnDestroy {
  private readonly _debounceTime = 500;
  readonly Controls = Controls;
  isPresetSelected = false;
  subnets: string[] = [];
  selectedSubnet = '';
  subnetLabel = SubnetState.Empty;
  vpcIDs: KubeVirtVPC[] = [];
  selectedVPC = '';
  vpcLabel = VPCState.Empty;

  @ViewChild('subnetCombobox')
  private readonly _subnetCombobox: FilteredComboboxComponent;
  @ViewChild('vpcCombobox')
  private readonly _vpcCombobox: FilteredComboboxComponent;

  private _kubeconfig = '';

  constructor(
    private readonly _cdr: ChangeDetectorRef,
    private readonly _builder: FormBuilder,
    private readonly _presets: PresetsService,
    private readonly _clusterSpecService: ClusterSpecService
  ) {
    super('KubeVirt Provider Extended');
  }

  ngOnInit(): void {
    this.form = this._builder.group({
      [Controls.Subnet]: this._builder.control(''),
      [Controls.VPCID]: this._builder.control(''),
    });

    this._presets.presetChanges.pipe(takeUntil(this._unsubscribe)).subscribe(preset => {
      this.isPresetSelected = !!preset;
      Object.values(Controls).forEach(control => {
        this._enable(!this.isPresetSelected, control);
      });
    });

    this.form.valueChanges
      .pipe(filter(_ => this._clusterSpecService.provider === NodeProvider.KUBEVIRT))
      .pipe(takeUntil(this._unsubscribe))
      .subscribe(_ =>
        this._presets.enablePresets(KubeVirtCloudSpec.isEmpty(this._clusterSpecService.cluster.spec.cloud.kubevirt))
      );

    merge(this._clusterSpecService.providerChanges, this._clusterSpecService.datacenterChanges)
      .pipe(takeUntil(this._unsubscribe))
      .subscribe(_ => this.form.reset());

    this._clusterSpecService.clusterChanges
      .pipe(filter(_ => this._clusterSpecService.provider === NodeProvider.KUBEVIRT))
      .pipe(debounceTime(this._debounceTime))
      .pipe(
        tap(_ => {
          if (!this._hasRequiredCredentials()) {
            this._clearSubnets();
            this._clearVPC();
          }
        })
      )
      .pipe(filter(_ => this._hasRequiredCredentials()))
      .pipe(filter(_ => this._areCredentialsChanged()))
      .pipe(tap(_ => this._vpcListObservable().subscribe(this._loadVPCs.bind(this))))
      .pipe(switchMap(_ => this._subnetObservable()))
      .pipe(takeUntil(this._unsubscribe))
      .subscribe(this._loadSubnets.bind(this));
  }

  ngOnDestroy(): void {
    this._unsubscribe.next();
    this._unsubscribe.complete();
  }

  getHint(control: Controls): string {
    switch (control) {
      case Controls.VPCID:
      case Controls.Subnet:
        return this._hasRequiredCredentials() ? '' : 'Please enter your credentials first.';
      default:
        return '';
    }
  }

  onVPCChange(vpcID: string): void {
    this.selectedVPC = vpcID;
    this._clusterSpecService.cluster.spec.cloud.kubevirt.vpcID = vpcID;
  }

  onSubnetChange(subnet: string): void {
    this.selectedSubnet = subnet;
    this._clusterSpecService.cluster.spec.cloud.kubevirt.subnet = subnet;
  }

  private _hasRequiredCredentials(): boolean {
    return !!this._clusterSpecService.cluster.spec.cloud.kubevirt?.kubeconfig;
  }

  private _areCredentialsChanged(): boolean {
    let credentialsChanged = false;

    if (this._clusterSpecService.cluster.spec.cloud.kubevirt?.kubeconfig !== this._kubeconfig) {
      this._kubeconfig = this._clusterSpecService.cluster.spec.cloud.kubevirt.kubeconfig;
      credentialsChanged = true;
    }
    return credentialsChanged;
  }

  private _vpcListObservable(): Observable<KubeVirtVPC[]> {
    return this._presets
      .provider(NodeProvider.KUBEVIRT)
      .kubeconfig(this._clusterSpecService.cluster.spec.cloud.kubevirt.kubeconfig)
      .vpcs(this._onVPCLoading.bind(this))
      .pipe(map(vpcs => _.sortBy(vpcs, v => v.name.toLowerCase())))
      .pipe(
        catchError(() => {
          this._clearVPC();
          return onErrorResumeNext(EMPTY);
        })
      );
  }

  private _onVPCLoading(): void {
    this._clearVPC();
    this.vpcLabel = VPCState.Loading;
    this._cdr.detectChanges();
  }

  private _clearVPC(): void {
    this.vpcIDs = [];
    this.selectedVPC = '';
    this.vpcLabel = VPCState.Empty;
    this._vpcCombobox.reset();
    this._cdr.detectChanges();
  }

  private _loadVPCs(vpcs: KubeVirtVPC[]): void {
    this.vpcIDs = vpcs;
    if (this.selectedVPC && !vpcs.find(vpc => vpc.vpcId === this.selectedVPC)) {
      this.selectedVPC = '';
    }
    if (!this.selectedVPC && vpcs.length) {
      const defaultVPC = this.vpcIDs.find(vpc => vpc.isDefault);
      this.selectedVPC = defaultVPC ? defaultVPC.vpcId : undefined;
    }
    this.vpcLabel = !_.isEmpty(this.vpcIDs) ? VPCState.Ready : VPCState.Empty;
    this._cdr.detectChanges();
  }

  private _subnetObservable(): Observable<string[]> {
    return this._presets
      .provider(NodeProvider.KUBEVIRT)
      .kubeconfig(this._clusterSpecService.cluster.spec.cloud.kubevirt.kubeconfig)
      .subnets(() => this._onSubnetsLoading.bind(this))
      .pipe(
        map(subnets => _.sortBy(subnets, s => s.toLowerCase())),
        catchError(() => {
          this._clearSubnets();
          return onErrorResumeNext(EMPTY);
        })
      );
  }

  private _clearSubnets(): void {
    this.subnets = [];
    this.subnetLabel = SubnetState.Empty;
    this._subnetCombobox.reset();
    this._cdr.detectChanges();
  }

  private _onSubnetsLoading(): void {
    this._clearSubnets();
    this.subnetLabel = SubnetState.Loading;
    this._cdr.detectChanges();
  }

  private _loadSubnets(subnets: string[]): void {
    this.subnets = subnets;
    this.subnetLabel = !_.isEmpty(this.subnets) ? SubnetState.Ready : SubnetState.Empty;
    this._cdr.detectChanges();
  }

  private _enable(enable: boolean, name: string): void {
    if (enable && this.form.get(name).disabled) {
      this.form.get(name).enable();
    }

    if (!enable && this.form.get(name).enabled) {
      this.form.get(name).disable();
    }
  }
}
