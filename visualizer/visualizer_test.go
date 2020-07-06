package visualizer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/genomagic/reporter"
)

func TestVisualize_Process(t *testing.T) {
	tests := []struct {
		testName    string
		vizName     string
		reports     []reporter.Report
		expectedErr error
	}{
		{
			testName: "test_returns_error_on_unprocessed_report",
			vizName:  "test1.html",
			reports: []reporter.Report{
				{
					AssemblyName: "Assembly 1",
					Processed:    false,
					N50:          0,
					L50:          0,
				},
			},
			expectedErr: fmt.Errorf("the reporter process has not been executed"),
		},
		{
			testName:    "test_returns_error_on_nil_reports",
			vizName:     "test2.html",
			reports:     nil,
			expectedErr: fmt.Errorf("visualizer has no reports to visualize"),
		},
		{
			testName: "test_visualizes_single_report",
			vizName:  "test3.html",
			reports: []reporter.Report{
				{
					AssemblyName: "Assembly 1",
					Processed:    true,
					N50:          10,
					L50:          11,
				},
			},
			expectedErr: nil,
		},
		{
			testName: "test_visualizes_multiple_reports",
			vizName:  "test4.html",
			reports: []reporter.Report{
				{
					AssemblyName: "Assembly 1",
					Processed:    true,
					N50:          10,
					L50:          11,
				},
				{
					AssemblyName: "Assembly 2",
					Processed:    true,
					N50:          15,
					L50:          16,
				},
				{
					AssemblyName: "Assembly 3",
					Processed:    true,
					N50:          20,
					L50:          21,
				},
			},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			viz := NewVisualizer(tt.reports, tt.vizName)
			err := viz.Process()
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}