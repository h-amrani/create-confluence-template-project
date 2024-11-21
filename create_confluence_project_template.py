import requests

# Confluence API configuration
base_url = "https://<your-confluence-domain>/wiki/rest/api/content"
auth = ("<your-username>", "<your-api-token>")

# Function to create a Confluence page
def create_page(title, parent_id, body, space_key="<space-key>"):
    data = {
        "type": "page",
        "title": title,
        "ancestors": [{"id": parent_id}] if parent_id else [],
        "space": {"key": space_key},
        "body": {"storage": {"value": body, "representation": "storage"}}
    }
    response = requests.post(base_url, json=data, auth=auth)
    if response.status_code in [200, 201]:
        print(f"Page '{title}' created successfully.")
        return response.json()["id"]
    else:
        print(f"Failed to create page '{title}': {response.status_code} - {response.text}")
        return None

# Create the Home Page
home_page_body = "<h1>Welcome to the Project</h1>" \
                 "<p>This space is dedicated to organizing the entire project lifecycle.</p>"
home_page_id = create_page("Project Name Overview", None, home_page_body)

# Define main sections with their subcategories
sections = [
    {
        "title": "Project Overview",
        "body": "<h1>Project Overview</h1><p>Objectives, scope, and assumptions for the project.</p>",
        "subcategories": [
            {"title": "Project Objectives", "body": "<p>High-level description of the project's goals.</p>"},
            {"title": "Scope", "body": "<p>In-scope and out-of-scope items.</p>"},
            {"title": "Assumptions", "body": "<p>Key dependencies and constraints.</p>"}
        ]
    },
    {
        "title": "Requirements",
        "body": "<h1>Requirements</h1><p>Functional and non-functional requirements, constraints, and workflows.</p>",
        "subcategories": [
            {"title": "Functional Requirements", "body": "<p>Core functionalities and user stories.</p>"},
            {"title": "Non-Functional Requirements", "body": "<p>Performance, scalability, and usability requirements.</p>"},
            {"title": "Constraints", "body": "<p>Technical and timeline limitations.</p>"},
            {"title": "AS-IS and TO-BE Processes", "body": "<p>Current and future workflows.</p>"}
        ]
    },
    {
        "title": "Architecture and Technical Design",
        "body": "<h1>Architecture and Technical Design</h1><p>System architecture and technical solutions.</p>",
        "subcategories": [
            {"title": "System Architecture", "body": "<p>High-level diagrams and descriptions.</p>"},
            {"title": "Data Flow Diagrams", "body": "<p>Data movement workflows.</p>"},
            {"title": "ETL Process Design", "body": "<p>ETL pipelines for data integration.</p>"},
            {"title": "Deployment Architecture", "body": "<p>Cloud infrastructure details.</p>"}
        ]
    },
    {
        "title": "Development Documentation",
        "body": "<h1>Development Documentation</h1><p>Technical guides for developers.</p>",
        "subcategories": [
            {"title": "API Documentation", "body": "<p>Endpoints and schemas.</p>"},
            {"title": "Database Schema", "body": "<p>Table design and indexing.</p>"},
            {"title": "CI/CD Pipeline", "body": "<p>Build, test, and deployment processes.</p>"}
        ]
    },
    {
        "title": "Testing and QA",
        "body": "<h1>Testing and QA</h1><p>Testing plans and criteria.</p>",
        "subcategories": [
            {"title": "Testing Strategy", "body": "<p>Unit, integration, and end-to-end tests.</p>"},
            {"title": "Test Cases", "body": "<p>Detailed test cases.</p>"},
            {"title": "UAT Plans", "body": "<p>User acceptance testing.</p>"}
        ]
    },
    {
        "title": "Reporting and Metrics",
        "body": "<h1>Reporting and Metrics</h1><p>Monitoring and reporting requirements.</p>",
        "subcategories": []
    },
    {
        "title": "Change Management",
        "body": "<h1>Change Management</h1><p>Track changes and approvals.</p>",
        "subcategories": []
    },
    {
        "title": "Training and Support",
        "body": "<h1>Training and Support</h1><p>User guides and support contacts.</p>",
        "subcategories": []
    }
]

# Create main sections and their subcategories
if home_page_id:
    for section in sections:
        section_id = create_page(section["title"], home_page_id, section["body"])
        if section_id:
            for sub in section.get("subcategories", []):
                create_page(sub["title"], section_id, sub["body"])