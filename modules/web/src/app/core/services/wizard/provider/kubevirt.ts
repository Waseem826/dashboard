// Copyright 2020 The Kubermatic Kubernetes Platform contributors.
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

import {HttpClient} from '@angular/common/http';
import {
  KubeVirtInstanceTypeList,
  KubeVirtOSImageList,
  KubeVirtPreferenceList,
  KubeVirtStorageClass,
  KubeVirtSubnets,
  KubeVirtVPC,
} from '@shared/entity/provider/kubevirt';
import {NodeProvider} from '@shared/model/NodeProviderConstants';
import {EMPTY, Observable} from 'rxjs';
import {map} from 'rxjs/operators';
import {Provider} from './provider';

export class KubeVirt extends Provider {
  constructor(http: HttpClient, projectID: string, provider: NodeProvider) {
    super(http, projectID, provider);

    this._setRequiredHeaders(KubeVirt.Header.Kubeconfig);
  }

  kubeconfig(token: string): KubeVirt {
    if (token) {
      this._headers = this._headers.set(KubeVirt.Header.Kubeconfig, token);
    }
    return this;
  }

  datacenterName(datacenterName: string): KubeVirt {
    if (datacenterName) {
      this._headers = this._headers.set(KubeVirt.Header.DatacenterName, datacenterName);
    }
    return this;
  }

  credential(credential: string): KubeVirt {
    super._credential(credential);
    return this;
  }

  vpcs(onLoadingCb: () => void = null): Observable<KubeVirtVPC[]> {
    if (!this._hasRequiredHeaders()) {
      return EMPTY;
    }

    if (onLoadingCb) {
      onLoadingCb();
    }

    const url = `${this._newRestRoot}/projects/${this._projectID}/providers/${this._provider}/vpcs`;
    return this._http
      .get<KubeVirtVPC[]>(url, {headers: this._headers})
      .pipe(map(vpcs => vpcs.map(vpc => Object.assign(new KubeVirtVPC(), vpc))));
  }

  subnets(onLoadingCb: () => void = null): Observable<string[]> {
    if (!this._hasRequiredHeaders()) {
      return EMPTY;
    }

    if (onLoadingCb) {
      onLoadingCb();
    }

    const url = `${this._newRestRoot}/projects/${this._projectID}/providers/${this._provider}/subnets`;
    return this._http.get<KubeVirtSubnets>(url, {headers: this._headers}).pipe(map(subnets => subnets.subnets));
  }

  instanceTypes(onLoadingCb: () => void = null): Observable<KubeVirtInstanceTypeList> {
    if (!this._hasRequiredHeaders()) {
      return EMPTY;
    }

    if (onLoadingCb) {
      onLoadingCb();
    }

    const url = `${this._newRestRoot}/projects/${this._projectID}/providers/${this._provider}/instancetypes`;
    return this._http.get<KubeVirtInstanceTypeList>(url, {headers: this._headers});
  }

  preferences(onLoadingCb: () => void = null): Observable<KubeVirtPreferenceList> {
    if (!this._hasRequiredHeaders()) {
      return EMPTY;
    }

    if (onLoadingCb) {
      onLoadingCb();
    }

    const url = `${this._newRestRoot}/projects/${this._projectID}/providers/${this._provider}/preferences`;
    return this._http.get<KubeVirtPreferenceList>(url, {headers: this._headers});
  }

  storageClass(onLoadingCb: () => void = null): Observable<KubeVirtStorageClass[]> {
    if (!this._hasRequiredHeaders()) {
      return EMPTY;
    }

    if (onLoadingCb) {
      onLoadingCb();
    }

    const url = `${this._newRestRoot}/projects/${this._projectID}/providers/${this._provider}/storageclasses`;
    return this._http.get<KubeVirtStorageClass[]>(url, {headers: this._headers});
  }

  osImages(dc: string, onLoadingCb: () => void = null): Observable<KubeVirtOSImageList> {
    if (!this._hasRequiredHeaders()) {
      return EMPTY;
    }

    if (onLoadingCb) {
      onLoadingCb();
    }
    const url = `${this._newRestRoot}/providers/${this._provider}/dc/${dc}/images`;
    return this._http.get<KubeVirtOSImageList>(url, {headers: this._headers});
  }
}

export namespace KubeVirt {
  export enum Header {
    Kubeconfig = 'Kubeconfig',
    DatacenterName = 'DatacenterName',
  }
}
