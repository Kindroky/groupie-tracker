# Groupie Tracker

**Groupie Tracker** is a web-based application that interacts with an API to display detailed information about various bands and artists. By leveraging the API, the project visualizes data about artists, their past and upcoming concerts, and how these elements are related. The project emphasizes both data manipulation and creating an interactive user experience through client-server communication.

## Project Overview

### API Structure
The API is divided into four key components:

1. **Artists**: Contains information about bands and artists, including:
   - Name(s)
   - Images
   - Year of activity initiation
   - Date of first album release
   - Band members

2. **Locations**: Provides the locations of the artists' past and upcoming concerts.

3. **Dates**: Lists the dates corresponding to the concert locations.

4. **Relation**: Links the artists to their concert locations and dates, acting as a bridge between the data points.

### Objective
The goal of Groupie Trackers is to transform the raw API data into a user-friendly and visually appealing website. The site will present band information in an easily digestible format using various data visualization techniques (e.g., blocks, cards, tables, lists, pages, or graphics). Additionally, the project includes implementing an event or action that triggers a client-server interaction, allowing for real-time data fetching or updates.

## Features

- **Data Visualization**: Display artists' information in intuitive ways such as cards, tables, or graphics.
- **Interactive User Interface**: Client-side interactions trigger server requests to fetch or manipulate data, providing a dynamic browsing experience.
- **Error Handling**: All pages are designed to function correctly, with robust error handling to ensure the site and server do not crash.
- **Backend in Go**: The backend logic is implemented in Go, utilizing only standard Go packages for data handling, JSON manipulation, and server requests.

## Technical Requirements

- **Backend**: Must be implemented in Go, following best practices for efficient, reliable code.
- **Frontend**: The website must offer a clean, user-friendly interface for displaying band-related data.
- **Client-Server Communication**: The project should include a feature that triggers an event, causing the server to respond to the client’s request. This could be a button click, time-triggered action, or other client-side events.
- **Testing**: It is recommended to include unit tests to ensure the accuracy and stability of the backend functionality.

## Key Learning Points

This project will help developers deepen their understanding of:

- **Data Manipulation**: Efficiently handling and organizing large sets of data for easy retrieval and display.
- **Client-Server Architecture**: Understanding the request-response cycle between clients and servers.
- **JSON Parsing**: Handling JSON data format for storage and display.
- **Event Creation**: Developing interactive features where user actions trigger backend processes.
- **HTML/CSS**: Building a responsive and user-friendly website for presenting data.

## Instructions for Use

To set up and run the project:

1. Clone the repository.
2. Ensure you have Go installed.
3. Run the Go server.
4. Open the website in a browser to view the visualizations and interact with the data.
