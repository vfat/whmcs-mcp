# WHMCS MCP Server (Go)

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://golang.org)
[![MCP Specification](https://img.shields.io/badge/MCP-2024--11--05-blue)](https://modelcontextprotocol.io)
[![Test Suite](https://img.shields.io/badge/tests-165%20passed-brightgreen)](https://github.com/vfat/whmcs-mcp)
[![Parity](https://img.shields.io/badge/parity-100%25%20(62%20Tools)-success)](https://github.com/vfat/whmcs-mcp)

A production-grade, high-performance Model Context Protocol (MCP) server written in Golang for managing WHMCS installations. Enables AI assistants (Claude Desktop, Cursor, Continue, etc.) to interact directly with clients, billing, support tickets, domains, orders, quotes, server provisioning modules, and administrative tasks.

---

## 🚀 Key Features & Highlights

- **100% Functional Parity**: Complete coverage matching the reference WHMCS MCP ecosystem:
  - **62 MCP Tools**: Complete administrative coverage (Clients, Billing, Invoices, Support, Domains, Orders, Quotes, Provisioning, Affiliates, Promotions).
  - **11 MCP Resources**: Instant read-only telemetry across stats, currencies, payment methods, active products, server nodes, departments, and TLD pricing schemes (`whmcs://*`).
  - **8 MCP Prompts**: Pre-engineered LLM workflows for onboarding, support ticket handling, revenue analysis, bulk overdue reminders, domain audit, fraud investigation, and new package setup.
- **Strict Stdio Protocol Isolation**: Logging is strictly redirected to `os.Stderr` via structured `log/slog`, guaranteeing that `os.Stdout` remains 100% pure JSON-RPC 2.0 without corruption.
- **High-Performance HTTP Engine**: Built-in Go connection pooling, request cancellation via context propagation, custom reflection serializer for PHP-style nested forms (`application/x-www-form-urlencoded`).
- **Test-Driven Design (TDD)**: Verified by 165+ automated test assertions with comprehensive unit and integration suites.

---

## 📦 Architecture & Inventory

### 1. Tools (62 Tools)

| Domain | Count | Highlighted Tools |
|---|:---:|---|
| **Client Management** | 9 | `whmcs_get_clients`, `whmcs_get_client_details`, `whmcs_add_client`, `whmcs_update_client`, `whmcs_delete_client`, `whmcs_get_client_products`, `whmcs_get_client_domains`, `whmcs_get_client_invoices`, `whmcs_close_client` |
| **Billing & Finance** | 9 | `whmcs_get_invoices`, `whmcs_get_invoice`, `whmcs_create_invoice`, `whmcs_update_invoice`, `whmcs_add_payment`, `whmcs_apply_credit`, `whmcs_get_transactions`, `whmcs_get_products`, `whmcs_get_product_groups` |
| **Support Tickets** | 9 | `whmcs_get_tickets`, `whmcs_get_ticket`, `whmcs_open_ticket`, `whmcs_add_ticket_reply`, `whmcs_add_ticket_note`, `whmcs_update_ticket`, `whmcs_delete_ticket`, `whmcs_get_support_departments`, `whmcs_get_support_statuses` |
| **Domain Lifecycle** | 9 | `whmcs_register_domain`, `whmcs_transfer_domain`, `whmcs_renew_domain`, `whmcs_get_domain_whois`, `whmcs_get_domain_nameservers`, `whmcs_update_domain_nameservers`, `whmcs_get_domain_lock_status`, `whmcs_update_domain_lock_status`, `whmcs_get_tld_pricing` |
| **Orders & Quotes** | 10 | `whmcs_get_orders`, `whmcs_accept_order`, `whmcs_cancel_order`, `whmcs_delete_order`, `whmcs_fraud_order`, `whmcs_pending_order`, `whmcs_get_quotes`, `whmcs_create_quote`, `whmcs_accept_quote`, `whmcs_delete_quote` |
| **Provisioning & Servers** | 6 | `whmcs_get_servers`, `whmcs_module_create`, `whmcs_module_suspend`, `whmcs_module_unsuspend`, `whmcs_module_terminate`, `whmcs_module_change_password` |
| **Marketing & Growth** | 3 | `whmcs_get_affiliates`, `whmcs_activate_affiliate`, `whmcs_get_promotions` |
| **System Administration** | 7 | `whmcs_get_stats`, `whmcs_get_admin_users`, `whmcs_get_payment_methods`, `whmcs_get_currencies`, `whmcs_get_activity_log`, `whmcs_log_activity`, `whmcs_get_email_templates`, `whmcs_send_email`, `whmcs_get_todo_items`, `whmcs_update_todo_item` |

### 2. Resources (11 URI Schemes)

- `whmcs://stats`: High-level operational statistics (income, ticket counts, active orders).
- `whmcs://admin-users`: Active administrative staff accounts.
- `whmcs://currencies`: Active billing currencies and conversion rates.
- `whmcs://payment-methods`: Configured payment gateways.
- `whmcs://admin/todo`: Pending administrative to-do tasks.
- `whmcs://products`: Active hosting packages and product catalog.
- `whmcs://support/departments`: Support departments and open ticket metrics.
- `whmcs://support/statuses`: Configured ticket statuses.
- `whmcs://tld-pricing`: Domain registration, transfer, and renewal pricing table.
- `whmcs://servers`: Registered server nodes and utilization metrics.
- `whmcs://promotions`: Active promotional codes and discounts.

### 3. Prompts (8 Interactive Workflows)

- `client-onboarding`: Guided workflow for creating client accounts, services, and welcome notifications.
- `client-health-check`: Client account audit (overdue invoices, tickets, active services).
- `ticket-response`: Intelligent support ticket response formulation with previous thread context.
- `revenue-report`: Periodic revenue and receivables analysis (daily, weekly, monthly, annual).
- `bulk-invoice-reminder`: Systematic overdue debt collection strategy and notification drafting.
- `domain-expiry-audit`: Audit of expiring domains, autorenew status, and registrar locks.
- `fraud-investigation`: Security assessment of suspicious orders (IP, location, order patterns).
- `new-product-setup`: Guided configuration for new hosting packages and server module settings.

---

## 🛠️ Build & Installation

### Requirements
- **Go**: 1.22 or newer.
- WHMCS installation with API credentials (API Role permissions enabled).

### Build from Source
```bash
git clone https://github.com/vfat/whmcs-mcp.git
cd whmcs-mcp
go build -o whmcs-mcp ./cmd/whmcs-mcp
```

### Run Tests
```bash
go test -v ./...
```

---

## ⚙️ Configuration

Configure the server via environment variables or a `.env` file:

| Variable | Required | Default | Description |
|---|:---:|:---:|---|
| `WHMCS_URL` | **Yes** | - | Base URL of your WHMCS installation (e.g. `https://billing.example.com/whmcs/`) |
| `WHMCS_API_IDENTIFIER` | **Yes** | - | WHMCS API Credential Identifier |
| `WHMCS_API_SECRET` | **Yes** | - | WHMCS API Credential Secret |
| `WHMCS_API_ACCESS_KEY` | No | - | Optional 2FA / API Access Key |
| `WHMCS_TIMEOUT` | No | `30s` | HTTP request timeout |
| `WHMCS_DEBUG` | No | `false` | Enable verbose debug logging to stderr |

*Legacy aliases `WHMCS_API_URL` and `WHMCS_ACCESS_KEY` are also automatically recognized.*

---

## 🔌 Client Integration

### Claude Desktop Configuration
Add the server to your `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "whmcs": {
      "command": "/path/to/whmcs-mcp",
      "env": {
        "WHMCS_URL": "https://billing.example.com/whmcs/",
        "WHMCS_API_IDENTIFIER": "YOUR_API_IDENTIFIER",
        "WHMCS_API_SECRET": "YOUR_API_SECRET",
        "WHMCS_API_ACCESS_KEY": "YOUR_OPTIONAL_ACCESS_KEY"
      }
    }
  }
}
```

### Cursor IDE Configuration
In your Cursor settings under **Features > MCP Servers**, add a new stdio transport:
- **Name**: `whmcs`
- **Command**: `/path/to/whmcs-mcp`

---

## 📄 License

MIT License. See [LICENSE](LICENSE) for details.
