package namespaces

import (
	"testing"
	"time"

	"github.com/stakater/Jamadar/internal/pkg/actions"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	testclient "k8s.io/client-go/kubernetes/fake"
)

func TestNamespaceToDelete_DeleteNamespaces(t *testing.T) {
	type fields struct {
		clientset            clientset.Interface
		actions              []actions.Action
		age                  string
		restrictedNamespaces []string
	}
	type args struct {
		namespaceList *v1.NamespaceList
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			// 1st has a less time so don't delete
			// 2nd will be deleted
			name: "DeleteNamespaces",
			args: args{
				namespaceList: &v1.NamespaceList{
					Items: []v1.Namespace{
						v1.Namespace{
							ObjectMeta: metav1.ObjectMeta{
								Name: "ns-test1",
								Annotations: map[string]string{
									"jamadar.stakater.com/persist": "false",
								},
								CreationTimestamp: metav1.NewTime(time.Now()),
							},
						},
						v1.Namespace{
							ObjectMeta: metav1.ObjectMeta{
								Name: "ns-test2",
								Annotations: map[string]string{
									"jamadar.stakater.com/persist": "false",
								},
							},
						},
					},
				},
			},
			fields: fields{
				clientset: testclient.NewSimpleClientset(),
				age:       "1m",
				actions: []actions.Action{
					&actions.Default{},
				},
			},
			wantErr: false,
		},
		{
			name: "DontDeleteNamespaces",
			args: args{
				namespaceList: &v1.NamespaceList{
					Items: []v1.Namespace{
						v1.Namespace{
							ObjectMeta: metav1.ObjectMeta{
								Name: "ns-test1",
								Annotations: map[string]string{
									"jamadar.stakater.com/persist": "false",
								},
							},
						},
						v1.Namespace{
							ObjectMeta: metav1.ObjectMeta{
								Name: "ns-test2",
								// Annotations: map[string]string{
								// 	"jamadar.stakater.com/persist": "false",
								// },
							},
						},
					},
				},
			},
			fields: fields{
				clientset: testclient.NewSimpleClientset(),
				age:       "1m",
				actions: []actions.Action{
					&actions.Default{},
				},
				restrictedNamespaces: []string{
					"ns-test1",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NamespaceToDelete{
				clientset:            tt.fields.clientset,
				actions:              tt.fields.actions,
				age:                  tt.fields.age,
				restrictedNamespaces: tt.fields.restrictedNamespaces,
			}
			tt.fields.clientset.CoreV1().Namespaces().Create(&tt.args.namespaceList.Items[0])
			tt.fields.clientset.CoreV1().Namespaces().Create(&tt.args.namespaceList.Items[1])
			if err := n.DeleteNamespaces(tt.args.namespaceList); (err != nil) != tt.wantErr {
				t.Errorf("NamespaceToDelete.DeleteNamespaces() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestDeleteNamespacesWithNoNamespaceCreatedError(t *testing.T) {
	namespaceList := &v1.NamespaceList{
		Items: []v1.Namespace{
			// Not Old
			v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "ns-test1",
					Annotations: map[string]string{
						"jamadar.stakater.com/persist": "false",
					},
					CreationTimestamp: metav1.NewTime(time.Now()),
				},
			},
			// Actual test with actions that passes
			v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "ns-test2",
					Annotations: map[string]string{
						"jamadar.stakater.com/persist": "false",
					},
				},
			},
			// Gives error as namespace not created
			v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "ns-test3",
					Annotations: map[string]string{
						"jamadar.stakater.com/persist": "false",
					},
				},
			},
		},
	}
	clientset := testclient.NewSimpleClientset()
	clientset.CoreV1().Namespaces().Create(&namespaceList.Items[1])
	wantErr := true
	actions := []actions.Action{
		&actions.Default{},
	}
	n := &NamespaceToDelete{
		clientset:            clientset,
		actions:              actions,
		age:                  "1m",
		restrictedNamespaces: []string{},
	}
	err := n.DeleteNamespaces(namespaceList)
	if (err != nil) != wantErr {
		t.Errorf("DeleteNamespaces() error = %v, wantErr %v", err, wantErr)
		return
	}
}
