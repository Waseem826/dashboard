// Copyright 2023 The Kubermatic Kubernetes Platform contributors.
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

import {ChangeDetectionStrategy, Component, forwardRef, OnDestroy, OnInit} from '@angular/core';
import {FormBuilder, NG_VALIDATORS, NG_VALUE_ACCESSOR, Validators} from '@angular/forms';
import {KubeOneClusterSpecService} from '@core/services/kubeone-cluster-spec';
import {KubeOnePresetsService} from '@core/services/kubeone-wizard/kubeone-presets';
import {ExternalCloudSpec, ExternalCluster} from '@shared/entity/external-cluster';
import {
  KubeOneCloudSpec,
  KubeOneClusterSpec,
  KubeOneVMwareCloudDirectorCloudSpec,
} from '@shared/entity/kubeone-cluster';
import {BaseFormValidator} from '@shared/validators/base-form.validator';
import {merge} from 'rxjs';
import {distinctUntilChanged, takeUntil} from 'rxjs/operators';

export enum Controls {
  Username = 'username',
  Password = 'password',
  Organization = 'organization',
  Url = 'url',
  Vdc = 'vdc',
}

@Component({
  selector: 'km-kubeone-wizard-vmware-cloud-director-credentials-basic',
  templateUrl: './template.html',
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => KubeOneVMwareCloudDirectorCredentialsBasicComponent),
      multi: true,
    },
    {
      provide: NG_VALIDATORS,
      useExisting: forwardRef(() => KubeOneVMwareCloudDirectorCredentialsBasicComponent),
      multi: true,
    },
  ],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class KubeOneVMwareCloudDirectorCredentialsBasicComponent
  extends BaseFormValidator
  implements OnInit, OnDestroy
{
  readonly Controls = Controls;

  constructor(
    private readonly _builder: FormBuilder,
    private readonly _clusterSpecService: KubeOneClusterSpecService,
    private readonly _presetsService: KubeOnePresetsService
  ) {
    super('VMware Cloud Director Credentials Basic');
  }

  ngOnInit(): void {
    this._initForm();
    this._initSubscriptions();
  }

  ngOnDestroy(): void {
    this._unsubscribe.next();
    this._unsubscribe.complete();
  }

  private _initForm(): void {
    this.form = this._builder.group({
      [Controls.Username]: this._builder.control('', [Validators.required]),
      [Controls.Password]: this._builder.control('', [Validators.required]),
      [Controls.Organization]: this._builder.control('', [Validators.required]),
      [Controls.Url]: this._builder.control('', [Validators.required]),
      [Controls.Vdc]: this._builder.control('', [Validators.required]),
    });
  }

  private _initSubscriptions(): void {
    this._clusterSpecService.providerChanges.pipe(takeUntil(this._unsubscribe)).subscribe(_ => this.form.reset());

    this.form.valueChanges
      .pipe(takeUntil(this._unsubscribe))
      .subscribe(_ =>
        this._presetsService.enablePresets(Object.values(Controls).every(control => !this.form.get(control).value))
      );

    this._presetsService.presetChanges
      .pipe(takeUntil(this._unsubscribe))
      .subscribe(preset => Object.values(Controls).forEach(control => this._enable(!preset, control)));

    merge(
      this.form.get(Controls.Username).valueChanges,
      this.form.get(Controls.Password).valueChanges,
      this.form.get(Controls.Organization).valueChanges,
      this.form.get(Controls.Url).valueChanges,
      this.form.get(Controls.Vdc).valueChanges
    )
      .pipe(distinctUntilChanged())
      .pipe(takeUntil(this._unsubscribe))
      .subscribe(_ => (this._clusterSpecService.cluster = this._getClusterEntity()));
  }

  private _enable(enable: boolean, name: Controls): void {
    if (enable && this.form.get(name).disabled) {
      this.form.get(name).enable();
    }

    if (!enable && this.form.get(name).enabled) {
      this.form.get(name).disable();
    }
  }

  private _getClusterEntity(): ExternalCluster {
    return {
      cloud: {
        kubeOne: {
          cloudSpec: {
            vmwareclouddirector: {
              username: this.form.get(Controls.Username).value,
              password: this.form.get(Controls.Password).value,
              organization: this.form.get(Controls.Organization).value,
              url: this.form.get(Controls.Url).value,
              vdc: this.form.get(Controls.Vdc).value,
            } as KubeOneVMwareCloudDirectorCloudSpec,
          } as KubeOneCloudSpec,
        } as KubeOneClusterSpec,
      } as ExternalCloudSpec,
    } as ExternalCluster;
  }
}
