package tools_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/vfat/whmcs-mcp/internal/config"
	mcpServer "github.com/vfat/whmcs-mcp/internal/mcp"
	"github.com/vfat/whmcs-mcp/internal/mcp/tools"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

func TestRegisterTicketTools_AllPresent(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	tools.RegisterTicketTools(srv.MCPServer(), client)

	expected := []string{
		"whmcs_get_tickets",
		"whmcs_get_ticket",
		"whmcs_open_ticket",
		"whmcs_add_ticket_reply",
		"whmcs_add_ticket_note",
		"whmcs_update_ticket",
		"whmcs_delete_ticket",
		"whmcs_get_support_departments",
		"whmcs_get_support_statuses",
	}

	toolMap := make(map[string]bool)
	for _, tl := range srv.ListTools() {
		toolMap[tl.Name] = true
	}

	for _, name := range expected {
		if !toolMap[name] {
			t.Errorf("expected tool %s to be registered", name)
		}
	}
}

func TestTicketTools_Execution(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		action := r.URL.Query().Get("action")
		if action == "" {
			action = r.FormValue("action")
		}

		switch action {
		case "GetTickets":
			_ = json.NewEncoder(w).Encode(model.GetTicketsResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   1,
			})
		case "GetTicket":
			_ = json.NewEncoder(w).Encode(model.GetTicketResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TicketID:       55,
			})
		case "OpenTicket":
			_ = json.NewEncoder(w).Encode(model.OpenTicketResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				ID:             56,
				TID:            "XYZ-9999",
			})
		case "AddTicketReply":
			_ = json.NewEncoder(w).Encode(model.AddTicketReplyResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
			})
		case "AddTicketNote":
			_ = json.NewEncoder(w).Encode(model.AddTicketNoteResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
			})
		case "UpdateTicket":
			_ = json.NewEncoder(w).Encode(model.UpdateTicketResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TicketID:       55,
			})
		case "DeleteTicket":
			_ = json.NewEncoder(w).Encode(model.DeleteTicketResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TicketID:       55,
			})
		case "GetSupportDepartments":
			_ = json.NewEncoder(w).Encode(model.GetSupportDepartmentsResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   2,
			})
		case "GetSupportStatuses":
			_ = json.NewEncoder(w).Encode(model.GetSupportStatusesResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   3,
			})
		default:
			_ = json.NewEncoder(w).Encode(map[string]string{"result": "error"})
		}
	}))
	defer mockServer.Close()

	cfg := &config.Config{
		URL:        mockServer.URL,
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	tools.RegisterTicketTools(srv.MCPServer(), client)

	tests := []struct {
		name      string
		toolName  string
		args      map[string]interface{}
		wantError bool
	}{
		{
			name:     "get tickets success",
			toolName: "whmcs_get_tickets",
			args: map[string]interface{}{
				"status": "Open",
			},
			wantError: false,
		},
		{
			name:      "get ticket missing ticketid",
			toolName:  "whmcs_get_ticket",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "get ticket success",
			toolName: "whmcs_get_ticket",
			args: map[string]interface{}{
				"ticketid": 55,
			},
			wantError: false,
		},
		{
			name:      "open ticket missing required fields",
			toolName:  "whmcs_open_ticket",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "open ticket success",
			toolName: "whmcs_open_ticket",
			args: map[string]interface{}{
				"deptid":   1,
				"subject":  "Help needed",
				"message":  "Please assist with my service.",
				"clientid": 42,
			},
			wantError: false,
		},
		{
			name:      "add reply missing ticketid",
			toolName:  "whmcs_add_ticket_reply",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "add reply success",
			toolName: "whmcs_add_ticket_reply",
			args: map[string]interface{}{
				"ticketid": 55,
				"message":  "We are looking into this.",
			},
			wantError: false,
		},
		{
			name:      "add note missing message",
			toolName:  "whmcs_add_ticket_note",
			args:      map[string]interface{}{"ticketid": 55},
			wantError: true,
		},
		{
			name:     "add note success",
			toolName: "whmcs_add_ticket_note",
			args: map[string]interface{}{
				"ticketid": 55,
				"message":  "Internal staff note here.",
			},
			wantError: false,
		},
		{
			name:      "update ticket missing ticketid",
			toolName:  "whmcs_update_ticket",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "update ticket success",
			toolName: "whmcs_update_ticket",
			args: map[string]interface{}{
				"ticketid": 55,
				"status":   "In Progress",
			},
			wantError: false,
		},
		{
			name:      "delete ticket missing ticketid",
			toolName:  "whmcs_delete_ticket",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "delete ticket success",
			toolName: "whmcs_delete_ticket",
			args: map[string]interface{}{
				"ticketid": 55,
			},
			wantError: false,
		},
		{
			name:      "get support departments success",
			toolName:  "whmcs_get_support_departments",
			args:      map[string]interface{}{},
			wantError: false,
		},
		{
			name:      "get support statuses success",
			toolName:  "whmcs_get_support_statuses",
			args:      map[string]interface{}{},
			wantError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			toolItem := srv.MCPServer().GetTool(tc.toolName)
			if toolItem == nil {
				t.Fatalf("tool %s not registered", tc.toolName)
			}
			req := mcp.CallToolRequest{}
			req.Params.Name = tc.toolName
			req.Params.Arguments = tc.args

			res, err := toolItem.Handler(context.Background(), req)
			if err != nil {
				t.Fatalf("unexpected handler err: %v", err)
			}
			if tc.wantError && !res.IsError {
				t.Errorf("expected error result but got success")
			}
			if !tc.wantError && res.IsError {
				t.Errorf("expected success result but got error: %v", res.Content)
			}
		})
	}
}
