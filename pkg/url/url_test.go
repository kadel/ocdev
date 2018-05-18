package url

import (
	"reflect"
	"testing"

	routev1 "github.com/openshift/api/route/v1"
	"github.com/redhat-developer/odo/pkg/occlient"
	"k8s.io/apimachinery/pkg/runtime"
	ktesting "k8s.io/client-go/testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		componentName   string
		applicationName string
	}
	tests := []struct {
		name    string
		args    args
		want    *URL
		wantErr bool
	}{
		{
			name: "first test",
			args: args{
				componentName:   "component",
				applicationName: "application",
			},
			want: &URL{
				Name:     "component",
				Protocol: "http",
				URL:      "host",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, fkClientSet := occlient.FakeNew()

			fkClientSet.RouteClientset.PrependReactor("create", "routes", func(action ktesting.Action) (bool, runtime.Object, error) {
				route := action.(ktesting.CreateAction).GetObject().(*routev1.Route)
				route.Spec.Host = "host"
				return true, route, nil
			})

			got, err := Create(client, tt.args.componentName, tt.args.applicationName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %#v, wantErr %#v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
