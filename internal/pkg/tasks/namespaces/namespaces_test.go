package namespaces

import (
	"testing"
	"time"

	"github.com/stakater/Jamadaar/internal/pkg/actions"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	testclient "k8s.io/client-go/kubernetes/fake"
)

func TestDeleteNamespace(t *testing.T) {
	type args struct {
		clientset clientset.Interface
		namespace v1.Namespace
		age       string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "DeleteNamespaceNoAnnotation",
			args: args{
				clientset: testclient.NewSimpleClientset(),
				namespace: v1.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name: "ns-test",
					},
				},
				age: "1d",
			},
			want: true,
		},
		{
			name: "DeleteNamespaceAnnotationFalse",
			args: args{
				clientset: testclient.NewSimpleClientset(),
				namespace: v1.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name: "ns-test",
						Annotations: map[string]string{
							"jamadaar.stakater.com/persist": "false",
						},
					},
				},
				age: "1w",
			},
			want: true,
		},
		{
			name: "DeleteNamespaceTimeFalse",
			args: args{
				clientset: testclient.NewSimpleClientset(),
				namespace: v1.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name: "ns-test",
						Annotations: map[string]string{
							"jamadaar.stakater.com/persist": "false",
						},
						CreationTimestamp: metav1.NewTime(time.Now()),
					},
				},
				age: "1m",
			},
			want: false,
		},
		{
			name: "DeleteNamespaceAnnotationTrue",
			args: args{
				clientset: testclient.NewSimpleClientset(),
				namespace: v1.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name: "ns-test",
						Annotations: map[string]string{
							"jamadaar.stakater.com/persist": "true",
						},
					},
				},
				age: "1d",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.clientset.CoreV1().Namespaces().Create(&tt.args.namespace)
			got, err := DeleteNamespace(tt.args.clientset, tt.args.namespace, tt.args.age)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteNamespace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestDeleteNamespaceWithNoNamespaceCreatedError(t *testing.T) {
	namespace := v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ns-test",
			Annotations: map[string]string{
				"jamadaar.stakater.com/persist": "false",
			},
		},
	}
	wantErr := true
	_, err := DeleteNamespace(testclient.NewSimpleClientset(), namespace, "1y")
	if (err != nil) != wantErr {
		t.Errorf("DeleteNamespace() error = %v, wantErr %v", err, wantErr)
		return
	}
}

func TestDeleteNamespaces(t *testing.T) {
	type args struct {
		clientset     clientset.Interface
		namespaceList *v1.NamespaceList
		actions       []actions.Action
		age           string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "DeleteNamespaces",
			args: args{
				clientset: testclient.NewSimpleClientset(),
				namespaceList: &v1.NamespaceList{
					Items: []v1.Namespace{
						v1.Namespace{
							ObjectMeta: metav1.ObjectMeta{
								Name: "ns-test1",
								Annotations: map[string]string{
									"jamadaar.stakater.com/persist": "false",
								},
								CreationTimestamp: metav1.NewTime(time.Now()),
							},
						},
						v1.Namespace{
							ObjectMeta: metav1.ObjectMeta{
								Name: "ns-test2",
								Annotations: map[string]string{
									"jamadaar.stakater.com/persist": "false",
								},
							},
						},
					},
				},
				age: "1m",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.clientset.CoreV1().Namespaces().Create(&tt.args.namespaceList.Items[0])
			tt.args.clientset.CoreV1().Namespaces().Create(&tt.args.namespaceList.Items[1])
			if err := DeleteNamespaces(tt.args.clientset, tt.args.namespaceList, tt.args.actions, tt.args.age); (err != nil) != tt.wantErr {
				t.Errorf("DeleteNamespaces() error = %v, wantErr %v", err, tt.wantErr)
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
						"jamadaar.stakater.com/persist": "false",
					},
					CreationTimestamp: metav1.NewTime(time.Now()),
				},
			},
			// Actual test with actions that passes
			v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "ns-test2",
					Annotations: map[string]string{
						"jamadaar.stakater.com/persist": "false",
					},
				},
			},
			// Gives error as namespace not created
			v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "ns-test3",
					Annotations: map[string]string{
						"jamadaar.stakater.com/persist": "false",
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
	err := DeleteNamespaces(clientset, namespaceList, actions, "1y")
	if (err != nil) != wantErr {
		t.Errorf("DeleteNamespaces() error = %v, wantErr %v", err, wantErr)
		return
	}
}
