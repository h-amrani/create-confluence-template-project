# Confluence Project Template Creator

This repository contains a Python script to automate the creation of structured Confluence spaces for new projects. The script dynamically generates a hierarchy of pages, ensuring consistency and saving time during project setup.

## Features

- Creates a **Home Page** with a high-level overview.
- Adds **Main Sections** such as:
  - Project Overview
  - Requirements
  - Architecture and Technical Design
  - Development Documentation
  - Testing and QA
  - Reporting and Metrics
  - Change Management
  - Training and Support
- Includes **Subcategories** for detailed documentation under each section (e.g., Project Objectives, Functional Requirements).
- Easily customizable to match specific project needs.

## Prerequisites

- **Python 3.7 or higher**.
- The `requests` library. Install it using:

  ```bash
  pip install requests
  ```
- A valid Confluence account with:
- Access to the Confluence REST API.
- Add Pages permission in the target Confluence space.
- A Confluence API token. Generate it from Atlassian API Tokens.

## Usage

Step 1: Clone the Repository
 ```bash 
git clone https://github.com/your-username/confluence-project-template-creator.git
cd confluence-project-template-creator
```

Step 2: Configure the Script

	1.	Open Create_Project_Template.py.
	2.	Replace the following placeholders:
	•	<your-confluence-domain>: Your Confluence domain.
	•	<your-username>: Your Confluence username or email.
	•	<your-api-token>: Your Confluence API token.

Step 3: Run the Script
```bash  
python Create_Project_Template.py 
```

Step 4: Verify the Structure

	•	Log in to your Confluence space.
	•	Navigate to the newly created Home Page to verify the structure.


## Customization

The script is fully customizable. You can:
	•	Add or modify main sections and subcategories.
	•	Update the content (e.g., descriptions, instructions) for each page.