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
