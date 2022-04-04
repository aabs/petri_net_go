package persistence

import (
	"aabs/petri_net_go/core"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func Test_CreateEvent(t *testing.T) {
	ev := AssetFoundDuringOnboardingEvent{
		RecipientId: "RecipientId",
		AssetId:     "AssetId",
	}
	sut, err := CreateEvent("1.0", "test", ev)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "test", sut.EventType)
	assert.Equal(t, "1.0", sut.Version)
	assert.Equal(t, "AssetId", sut.Payload.(AssetFoundDuringOnboardingEvent).AssetId)
	assert.Equal(t, "RecipientId", sut.Payload.(AssetFoundDuringOnboardingEvent).RecipientId)
}

func Test_SetMarking(t *testing.T) {
	type args struct {
		m       *core.Marking
		placeId int
		token   float64
	}
	tests := []struct {
		name    string
		args    args
		want    *core.Marking
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "hello",
			args: args{
				m: &core.Marking{
					Places: mat.NewVecDense(5, []float64{0, 0, 0, 0, 0}),
				},
				placeId: 0,
				token:   1.0,
			},
			want: &core.Marking{
				Places: mat.NewVecDense(5, []float64{1.0, 0, 0, 0, 0}),
			},
			wantErr: false,
		}, {
			name: "set middle place",
			args: args{
				m: &core.Marking{
					Places: mat.NewVecDense(5, []float64{0, 0, 0, 0, 0}),
				},
				placeId: 2,
				token:   1.0,
			},
			want: &core.Marking{
				Places: mat.NewVecDense(5, []float64{0, 0, 1.0, 0, 0}),
			},
			wantErr: false,
		}, {
			name: "try to set missing place",
			args: args{
				m: &core.Marking{
					Places: mat.NewVecDense(5, []float64{0, 0, 0, 0, 0}),
				},
				placeId: 21,
				token:   1.0,
			},
			want: &core.Marking{
				Places: mat.NewVecDense(5, []float64{0, 0, 0, 0, 0}),
			},
			wantErr: true,
		},{
			name: "try to set missing place with negative index",
			args: args{
				m: &core.Marking{
					Places: mat.NewVecDense(5, []float64{0, 0, 0, 0, 0}),
				},
				placeId: -23,
				token:   1.0,
			},
			want: &core.Marking{
				Places: mat.NewVecDense(5, []float64{0, 0, 0, 0, 0}),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetMarking(tt.args.m, tt.args.placeId, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetMarking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				assert.NotNil(t, err)
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMarking() = %v, want %v", got, tt.want)
			}
		})
	}
}
