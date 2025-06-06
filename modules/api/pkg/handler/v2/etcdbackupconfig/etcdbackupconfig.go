/*
Copyright 2021 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package etcdbackupconfig

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

	apiv1 "k8c.io/dashboard/v2/pkg/api/v1"
	apiv2 "k8c.io/dashboard/v2/pkg/api/v2"
	handlercommon "k8c.io/dashboard/v2/pkg/handler/common"
	"k8c.io/dashboard/v2/pkg/handler/middleware"
	"k8c.io/dashboard/v2/pkg/handler/v1/common"
	"k8c.io/dashboard/v2/pkg/handler/v2/cluster"
	"k8c.io/dashboard/v2/pkg/provider"
	kubermaticv1 "k8c.io/kubermatic/sdk/v2/apis/kubermatic/v1"
	utilerrors "k8c.io/kubermatic/v2/pkg/util/errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/reference"
)

const (
	automaticBackup = "automatic"
	snapshot        = "snapshot"
)

func CreateEndpoint(userInfoGetter provider.UserInfoGetter, projectProvider provider.ProjectProvider,
	privilegedProjectProvider provider.PrivilegedProjectProvider, settingsProvider provider.SettingsProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createEtcdBackupConfigReq)

		if err := IsEtcdBackupEnabled(ctx, settingsProvider); err != nil {
			return nil, err
		}

		c, err := handlercommon.GetCluster(ctx, projectProvider, privilegedProjectProvider, userInfoGetter, req.ProjectID, req.ClusterID, nil)
		if err != nil {
			return nil, err
		}

		ebc, err := convertAPIToInternalEtcdBackupConfig(req.Body.Name, &req.Body.Spec, c)
		if err != nil {
			return nil, err
		}

		// set projectID label
		ebc.Labels = map[string]string{
			kubermaticv1.ProjectIDLabelKey: req.ProjectID,
		}

		ebc, err = createEtcdBackupConfig(ctx, userInfoGetter, req.ProjectID, ebc)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		return convertInternalToAPIEtcdBackupConfig(ebc), nil
	}
}

// createEtcdBackupConfigReq represents a request for creating a cluster etcd backup configuration
// swagger:parameters createEtcdBackupConfig
type createEtcdBackupConfigReq struct {
	cluster.GetClusterReq
	// in: body
	Body ebcBody
}

type ebcBody struct {
	// Name of the etcd backup config
	Name string `json:"name"`
	// EtcdBackupConfigSpec Spec of the etcd backup config
	Spec apiv2.EtcdBackupConfigSpec `json:"spec"`
}

func DecodeCreateEtcdBackupConfigReq(c context.Context, r *http.Request) (interface{}, error) {
	var req createEtcdBackupConfigReq
	cr, err := cluster.DecodeGetClusterReq(c, r)
	if err != nil {
		return nil, err
	}
	req.GetClusterReq = cr.(cluster.GetClusterReq)

	if err = json.NewDecoder(r.Body).Decode(&req.Body); err != nil {
		return nil, err
	}
	return req, nil
}

func GetEndpoint(userInfoGetter provider.UserInfoGetter, projectProvider provider.ProjectProvider,
	privilegedProjectProvider provider.PrivilegedProjectProvider, settingsProvider provider.SettingsProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEtcdBackupConfigReq)

		if err := IsEtcdBackupEnabled(ctx, settingsProvider); err != nil {
			return nil, err
		}

		c, err := handlercommon.GetCluster(ctx, projectProvider, privilegedProjectProvider, userInfoGetter, req.ProjectID, req.ClusterID, nil)
		if err != nil {
			return nil, err
		}

		ebc, err := getEtcdBackupConfig(ctx, userInfoGetter, c, req.ProjectID, decodeEtcdBackupConfigID(req.EtcdBackupConfigID, req.ClusterID))
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		return convertInternalToAPIEtcdBackupConfig(ebc), nil
	}
}

// getEtcdBackupConfigReq represents a request for getting a cluster etcd backup configuration
// swagger:parameters getEtcdBackupConfig deleteEtcdBackupConfig
type getEtcdBackupConfigReq struct {
	cluster.GetClusterReq
	// in: path
	// required: true
	EtcdBackupConfigID string `json:"ebc_id"`
}

func ListEndpoint(userInfoGetter provider.UserInfoGetter, projectProvider provider.ProjectProvider,
	privilegedProjectProvider provider.PrivilegedProjectProvider, settingsProvider provider.SettingsProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(listEtcdBackupConfigReq)

		if err := IsEtcdBackupEnabled(ctx, settingsProvider); err != nil {
			return nil, err
		}

		c, err := handlercommon.GetCluster(ctx, projectProvider, privilegedProjectProvider, userInfoGetter, req.ProjectID, req.ClusterID, nil)
		if err != nil {
			return nil, err
		}

		ebcList, err := listEtcdBackupConfig(ctx, userInfoGetter, c, req.ProjectID)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		var ebcAPIList []*apiv2.EtcdBackupConfig
		for _, ebc := range ebcList.Items {
			ebcAPIList = append(ebcAPIList, convertInternalToAPIEtcdBackupConfig(&ebc))
		}

		return ebcAPIList, nil
	}
}

// listEtcdBackupConfigReq represents a request for listing cluster etcd backup configurations
// swagger:parameters listEtcdBackupConfig
type listEtcdBackupConfigReq struct {
	cluster.GetClusterReq
}

func DecodeListEtcdBackupConfigReq(c context.Context, r *http.Request) (interface{}, error) {
	var req listEtcdBackupConfigReq

	cr, err := cluster.DecodeGetClusterReq(c, r)
	if err != nil {
		return nil, err
	}

	req.GetClusterReq = cr.(cluster.GetClusterReq)
	return req, nil
}

func DecodeGetEtcdBackupConfigReq(c context.Context, r *http.Request) (interface{}, error) {
	var req getEtcdBackupConfigReq
	cr, err := cluster.DecodeGetClusterReq(c, r)
	if err != nil {
		return nil, err
	}
	req.GetClusterReq = cr.(cluster.GetClusterReq)

	req.EtcdBackupConfigID = mux.Vars(r)["ebc_id"]
	if req.EtcdBackupConfigID == "" {
		return "", fmt.Errorf("'ebc_id' parameter is required but was not provided")
	}

	return req, nil
}

func DeleteEndpoint(userInfoGetter provider.UserInfoGetter, projectProvider provider.ProjectProvider,
	privilegedProjectProvider provider.PrivilegedProjectProvider, settingsProvider provider.SettingsProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEtcdBackupConfigReq)

		if err := IsEtcdBackupEnabled(ctx, settingsProvider); err != nil {
			return nil, err
		}

		c, err := handlercommon.GetCluster(ctx, projectProvider, privilegedProjectProvider, userInfoGetter, req.ProjectID, req.ClusterID, nil)
		if err != nil {
			return nil, err
		}

		err = deleteEtcdBackupConfig(ctx, userInfoGetter, c, req.ProjectID, decodeEtcdBackupConfigID(req.EtcdBackupConfigID, req.ClusterID))
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}
		return nil, nil
	}
}

func PatchEndpoint(userInfoGetter provider.UserInfoGetter, projectProvider provider.ProjectProvider,
	privilegedProjectProvider provider.PrivilegedProjectProvider, settingsProvider provider.SettingsProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(patchEtcdBackupConfigReq)

		if err := IsEtcdBackupEnabled(ctx, settingsProvider); err != nil {
			return nil, err
		}

		c, err := handlercommon.GetCluster(ctx, projectProvider, privilegedProjectProvider, userInfoGetter, req.ProjectID, req.ClusterID, nil)
		if err != nil {
			return nil, err
		}

		// get EBC
		originalEBC, err := getEtcdBackupConfig(ctx, userInfoGetter, c, req.ProjectID, decodeEtcdBackupConfigID(req.EtcdBackupConfigID, req.ClusterID))
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}
		newEBC := originalEBC.DeepCopy()
		newEBC.Spec.Keep = req.Body.Keep
		newEBC.Spec.Schedule = req.Body.Schedule
		newEBC.Spec.Destination = req.Body.Destination

		// apply patch
		ebc, err := patchEtcdBackupConfig(ctx, userInfoGetter, req.ProjectID, originalEBC, newEBC)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		return convertInternalToAPIEtcdBackupConfig(ebc), nil
	}
}

// patchEtcdBackupConfigReq represents a request for patching cluster etcd backup configurations
// swagger:parameters patchEtcdBackupConfig
type patchEtcdBackupConfigReq struct {
	cluster.GetClusterReq
	// in: path
	// required: true
	EtcdBackupConfigID string `json:"ebc_id"`
	// in: body
	// required: true
	Body apiv2.EtcdBackupConfigSpec
}

func DecodePatchEtcdBackupConfigReq(c context.Context, r *http.Request) (interface{}, error) {
	var req patchEtcdBackupConfigReq
	cr, err := cluster.DecodeGetClusterReq(c, r)
	if err != nil {
		return nil, err
	}
	req.GetClusterReq = cr.(cluster.GetClusterReq)

	req.EtcdBackupConfigID = mux.Vars(r)["ebc_id"]
	if req.EtcdBackupConfigID == "" {
		return "", fmt.Errorf("'ebc_id' parameter is required but was not provided")
	}

	if err = json.NewDecoder(r.Body).Decode(&req.Body); err != nil {
		return nil, err
	}

	return req, nil
}

func ProjectListEndpoint(userInfoGetter provider.UserInfoGetter, projectProvider provider.ProjectProvider,
	privilegedProjectProvider provider.PrivilegedProjectProvider, settingsProvider provider.SettingsProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(listProjectEtcdBackupConfigReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		// check if user has access to the project
		_, err := common.GetProject(ctx, userInfoGetter, projectProvider, privilegedProjectProvider, req.ProjectID, nil)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		if err := IsEtcdBackupEnabled(ctx, settingsProvider); err != nil {
			return nil, err
		}
		ebcLists, err := listProjectEtcdBackupConfig(ctx, req.ProjectID)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		var ebcAPIList []*apiv2.EtcdBackupConfig
		for _, ebcList := range ebcLists {
			for _, ebc := range ebcList.Items {
				ebcAPI := convertInternalToAPIEtcdBackupConfig(&ebc)
				if req.Type == "" || strings.EqualFold(req.Type, getEtcdBackupConfigType(ebcAPI)) {
					ebcAPIList = append(ebcAPIList, ebcAPI)
				}
			}
		}

		return ebcAPIList, nil
	}
}

func getEtcdBackupConfigType(ebc *apiv2.EtcdBackupConfig) string {
	if ebc.Spec.Schedule == "" {
		return snapshot
	}
	return automaticBackup
}

// listProjectEtcdBackupConfigReq represents a request for listing project etcd backupConfigs
// swagger:parameters listProjectEtcdBackupConfig
type listProjectEtcdBackupConfigReq struct {
	common.ProjectReq

	// in: query
	Type string `json:"type,omitempty"`
}

func (r *listProjectEtcdBackupConfigReq) validate() error {
	if len(r.Type) > 0 {
		if r.Type == automaticBackup || r.Type == snapshot {
			return nil
		}
		return utilerrors.NewBadRequest("wrong query parameter, unsupported type: %s", r.Type)
	}
	return nil
}

func DecodeListProjectEtcdBackupConfigReq(c context.Context, r *http.Request) (interface{}, error) {
	var req listProjectEtcdBackupConfigReq

	pr, err := common.DecodeProjectRequest(c, r)
	if err != nil {
		return nil, err
	}

	req.ProjectReq = pr.(common.ProjectReq)

	req.Type = r.URL.Query().Get("type")
	return req, nil
}

func convertInternalToAPIEtcdBackupConfig(ebc *kubermaticv1.EtcdBackupConfig) *apiv2.EtcdBackupConfig {
	etcdBackupConfig := &apiv2.EtcdBackupConfig{
		ObjectMeta: apiv1.ObjectMeta{
			Name:              ebc.Spec.Name,
			ID:                GenEtcdBackupConfigID(ebc.Name, ebc.Spec.Cluster.Name),
			Annotations:       ebc.Annotations,
			CreationTimestamp: apiv1.NewTime(ebc.CreationTimestamp.Time),
			DeletionTimestamp: func() *apiv1.Time {
				if ebc.DeletionTimestamp != nil {
					deletionTimestamp := apiv1.NewTime(ebc.DeletionTimestamp.Time)
					return &deletionTimestamp
				}
				return nil
			}(),
		},
		Spec: apiv2.EtcdBackupConfigSpec{
			ClusterID:   ebc.Spec.Cluster.Name,
			Schedule:    ebc.Spec.Schedule,
			Keep:        ebc.Spec.Keep,
			Destination: ebc.Spec.Destination,
		},
		Status: apiv2.EtcdBackupConfigStatus{
			CurrentBackups: []apiv2.BackupStatus{},
			Conditions:     []apiv2.EtcdBackupConfigCondition{},
			CleanupRunning: ebc.Status.CleanupRunning,
		},
	}

	for _, backupStatus := range ebc.Status.CurrentBackups {
		scheduledTime := apiv1.Time{}
		backupStartTime := apiv1.Time{}
		backupFinishedTime := apiv1.Time{}
		deleteStartTime := apiv1.Time{}
		deleteFinishedTime := apiv1.Time{}
		if !backupStatus.ScheduledTime.IsZero() {
			scheduledTime = apiv1.NewTime(backupStatus.ScheduledTime.Time)
		}
		if !backupStatus.BackupStartTime.IsZero() {
			backupStartTime = apiv1.NewTime(backupStatus.BackupStartTime.Time)
		}
		if !backupStatus.BackupFinishedTime.IsZero() {
			backupFinishedTime = apiv1.NewTime(backupStatus.BackupFinishedTime.Time)
		}
		if !backupStatus.DeleteStartTime.IsZero() {
			deleteStartTime = apiv1.NewTime(backupStatus.DeleteStartTime.Time)
		}
		if !backupStatus.DeleteFinishedTime.IsZero() {
			deleteFinishedTime = apiv1.NewTime(backupStatus.DeleteFinishedTime.Time)
		}

		apiBackupStatus := apiv2.BackupStatus{
			ScheduledTime:      &scheduledTime,
			BackupName:         backupStatus.BackupName,
			JobName:            backupStatus.JobName,
			BackupStartTime:    &backupStartTime,
			BackupFinishedTime: &backupFinishedTime,
			BackupPhase:        backupStatus.BackupPhase,
			BackupMessage:      backupStatus.BackupMessage,
			DeleteJobName:      backupStatus.DeleteJobName,
			DeleteStartTime:    &deleteStartTime,
			DeleteFinishedTime: &deleteFinishedTime,
			DeletePhase:        backupStatus.DeletePhase,
			DeleteMessage:      backupStatus.DeleteMessage,
		}
		etcdBackupConfig.Status.CurrentBackups = append(etcdBackupConfig.Status.CurrentBackups, apiBackupStatus)
	}

	for conditionType, condition := range ebc.Status.Conditions {
		lastHeartbeatTime := apiv1.NewTime(condition.LastHeartbeatTime.Time)
		lastTransitionTime := apiv1.NewTime(condition.LastTransitionTime.Time)

		apiCondition := apiv2.EtcdBackupConfigCondition{
			Type:               conditionType,
			Status:             condition.Status,
			LastHeartbeatTime:  lastHeartbeatTime,
			LastTransitionTime: lastTransitionTime,
			Reason:             condition.Reason,
			Message:            condition.Message,
		}
		etcdBackupConfig.Status.Conditions = append(etcdBackupConfig.Status.Conditions, apiCondition)
	}

	// ensure a stable sorting order
	sort.Slice(etcdBackupConfig.Status.Conditions, func(i, j int) bool {
		return etcdBackupConfig.Status.Conditions[i].Type < etcdBackupConfig.Status.Conditions[j].Type
	})

	return etcdBackupConfig
}

func convertAPIToInternalEtcdBackupConfig(name string, ebcSpec *apiv2.EtcdBackupConfigSpec, cluster *kubermaticv1.Cluster) (*kubermaticv1.EtcdBackupConfig, error) {
	clusterObjectRef, err := reference.GetReference(scheme.Scheme, cluster)
	if err != nil {
		return nil, utilerrors.New(http.StatusInternalServerError, fmt.Sprintf("error getting cluster object reference: %v", err))
	}

	return &kubermaticv1.EtcdBackupConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      rand.String(10),
			Namespace: cluster.Status.NamespaceName,
		},
		Spec: kubermaticv1.EtcdBackupConfigSpec{
			Name:        name,
			Cluster:     *clusterObjectRef,
			Schedule:    ebcSpec.Schedule,
			Keep:        ebcSpec.Keep,
			Destination: ebcSpec.Destination,
		},
	}, nil
}

func GenEtcdBackupConfigID(ebcName, clusterName string) string {
	return fmt.Sprintf("%s-%s", clusterName, ebcName)
}

// This is a little hack to avoid changing the EtcdBackupConfig API, as ebc ID is a parameter there, instead of the name.
// We can't have ID=name because ebc`s are also used in project wide LIST requests, so ID's need to be unique.
// Requests with just the Name will still work as the TripPrefix just won't remove anything.
func decodeEtcdBackupConfigID(id, clusterName string) string {
	return strings.TrimPrefix(id, fmt.Sprintf("%s-", clusterName))
}

func createEtcdBackupConfig(ctx context.Context, userInfoGetter provider.UserInfoGetter, projectID string, etcdBackupConfig *kubermaticv1.EtcdBackupConfig) (*kubermaticv1.EtcdBackupConfig, error) {
	adminUserInfo, privilegedEtcdBackupConfigProvider, err := getAdminUserInfoPrivilegedEtcdBackupConfigProvider(ctx, userInfoGetter)
	if err != nil {
		return nil, err
	}
	if adminUserInfo.IsAdmin {
		return privilegedEtcdBackupConfigProvider.CreateUnsecured(ctx, etcdBackupConfig)
	}
	userInfo, etcdBackupConfigProvider, err := getUserInfoEtcdBackupConfigProvider(ctx, userInfoGetter, projectID)
	if err != nil {
		return nil, err
	}
	return etcdBackupConfigProvider.Create(ctx, userInfo, etcdBackupConfig)
}

func getEtcdBackupConfig(ctx context.Context, userInfoGetter provider.UserInfoGetter, cluster *kubermaticv1.Cluster, projectID, etcdBackupConfigName string) (*kubermaticv1.EtcdBackupConfig, error) {
	adminUserInfo, privilegedEtcdBackupConfigProvider, err := getAdminUserInfoPrivilegedEtcdBackupConfigProvider(ctx, userInfoGetter)
	if err != nil {
		return nil, err
	}
	if adminUserInfo.IsAdmin {
		return privilegedEtcdBackupConfigProvider.GetUnsecured(ctx, cluster, etcdBackupConfigName)
	}
	userInfo, etcdBackupConfigProvider, err := getUserInfoEtcdBackupConfigProvider(ctx, userInfoGetter, projectID)
	if err != nil {
		return nil, err
	}
	return etcdBackupConfigProvider.Get(ctx, userInfo, cluster, etcdBackupConfigName)
}

func listEtcdBackupConfig(ctx context.Context, userInfoGetter provider.UserInfoGetter, cluster *kubermaticv1.Cluster, projectID string) (*kubermaticv1.EtcdBackupConfigList, error) {
	adminUserInfo, privilegedEtcdBackupConfigProvider, err := getAdminUserInfoPrivilegedEtcdBackupConfigProvider(ctx, userInfoGetter)
	if err != nil {
		return nil, err
	}
	if adminUserInfo.IsAdmin {
		return privilegedEtcdBackupConfigProvider.ListUnsecured(ctx, cluster)
	}
	userInfo, etcdBackupConfigProvider, err := getUserInfoEtcdBackupConfigProvider(ctx, userInfoGetter, projectID)
	if err != nil {
		return nil, err
	}
	return etcdBackupConfigProvider.List(ctx, userInfo, cluster)
}

func listProjectEtcdBackupConfig(ctx context.Context, projectID string) ([]*kubermaticv1.EtcdBackupConfigList, error) {
	privilegedEtcdBackupConfigProjectProvider := ctx.Value(middleware.PrivilegedEtcdBackupConfigProjectProviderContextKey).(provider.PrivilegedEtcdBackupConfigProjectProvider)
	if privilegedEtcdBackupConfigProjectProvider == nil {
		return nil, utilerrors.New(http.StatusInternalServerError, "error getting privileged provider")
	}
	return privilegedEtcdBackupConfigProjectProvider.ListUnsecured(ctx, projectID)
}

func deleteEtcdBackupConfig(ctx context.Context, userInfoGetter provider.UserInfoGetter, cluster *kubermaticv1.Cluster, projectID, etcdBackupConfigName string) error {
	adminUserInfo, privilegedEtcdBackupConfigProvider, err := getAdminUserInfoPrivilegedEtcdBackupConfigProvider(ctx, userInfoGetter)
	if err != nil {
		return err
	}
	if adminUserInfo.IsAdmin {
		return privilegedEtcdBackupConfigProvider.DeleteUnsecured(ctx, cluster, etcdBackupConfigName)
	}
	userInfo, etcdBackupConfigProvider, err := getUserInfoEtcdBackupConfigProvider(ctx, userInfoGetter, projectID)
	if err != nil {
		return err
	}
	return etcdBackupConfigProvider.Delete(ctx, userInfo, cluster, etcdBackupConfigName)
}

func patchEtcdBackupConfig(ctx context.Context, userInfoGetter provider.UserInfoGetter, projectID string, oldConfig, newConfig *kubermaticv1.EtcdBackupConfig) (*kubermaticv1.EtcdBackupConfig, error) {
	adminUserInfo, privilegedEtcdBackupConfigProvider, err := getAdminUserInfoPrivilegedEtcdBackupConfigProvider(ctx, userInfoGetter)
	if err != nil {
		return nil, err
	}
	if adminUserInfo.IsAdmin {
		return privilegedEtcdBackupConfigProvider.PatchUnsecured(ctx, oldConfig, newConfig)
	}
	userInfo, etcdBackupConfigProvider, err := getUserInfoEtcdBackupConfigProvider(ctx, userInfoGetter, projectID)
	if err != nil {
		return nil, err
	}
	return etcdBackupConfigProvider.Patch(ctx, userInfo, oldConfig, newConfig)
}

func getAdminUserInfoPrivilegedEtcdBackupConfigProvider(ctx context.Context, userInfoGetter provider.UserInfoGetter) (*provider.UserInfo, provider.PrivilegedEtcdBackupConfigProvider, error) {
	userInfo, err := userInfoGetter(ctx, "")
	if err != nil {
		return nil, nil, err
	}
	if !userInfo.IsAdmin {
		return userInfo, nil, nil
	}
	privilegedEtcdBackupConfigProvider := ctx.Value(middleware.PrivilegedEtcdBackupConfigProviderContextKey).(provider.PrivilegedEtcdBackupConfigProvider)
	return userInfo, privilegedEtcdBackupConfigProvider, nil
}

func getUserInfoEtcdBackupConfigProvider(ctx context.Context, userInfoGetter provider.UserInfoGetter, projectID string) (*provider.UserInfo, provider.EtcdBackupConfigProvider, error) {
	userInfo, err := userInfoGetter(ctx, projectID)
	if err != nil {
		return nil, nil, err
	}

	etcdBackupConfigProvider := ctx.Value(middleware.EtcdBackupConfigProviderContextKey).(provider.EtcdBackupConfigProvider)
	return userInfo, etcdBackupConfigProvider, nil
}

func IsEtcdBackupEnabled(ctx context.Context, settingsProvider provider.SettingsProvider) error {
	globalSettings, err := settingsProvider.GetGlobalSettings(ctx)

	if err != nil {
		return common.KubernetesErrorToHTTPError(err)
	}

	if !globalSettings.Spec.EnableEtcdBackup {
		return fmt.Errorf("etcd backup feature is disabled by the admin")
	}

	return nil
}
