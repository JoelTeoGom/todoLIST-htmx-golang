# HTMX & Go - Demo

## Description

This project is a web application built with Go (Golang) and HTMX that allows users to view a list of films, filter the list by title or director, and add new films to the list. The application demonstrates the use of HTMX to enhance the user experience by enabling dynamic interactions with the server without requiring full page reloads.

## Key Features

1. **Film List Display**: The application displays a list of films, each showing the title and the director. This list is generated on the server-side and rendered in the HTML template.

2. **Filter Films**: Users can filter the list of films by typing a title or director's name in the input box. The filtering is handled by the server, which responds with a filtered list, and HTMX updates the film list dynamically without refreshing the entire page.

3. **Add New Films**: Users can add new films to the list by submitting a form with the film's title and director. The server processes the form data, adds the new film to the list, and updates the displayed film list dynamically using HTMX.

4. **Responsive and Styled with Bootstrap**: The application uses Bootstrap 5 for styling, ensuring a responsive and visually appealing UI.

## Project Structure

- **index.html**: The main HTML file that contains the structure of the page, including the film list and the form for adding new films. This file utilizes Go's `html/template` package to dynamically render content based on the server-side data.

- **main.go**: The Go application code. It handles three main functionalities:
  - Serving the main page with the list of films.
  - Handling the addition of new films.
  - Filtering the list of films based on user input.

## How It Works

1. **HTMX Integration**: The `index.html` file includes the HTMX library, which is used to send HTTP requests triggered by user actions (e.g., typing in the filter input or submitting the add film form). HTMX handles these requests asynchronously and updates parts of the page (like the film list) based on the server's response.

2. **Server-Side Logic in Go**:
   - The `main.go` file sets up three handler functions for the routes `/`, `/add-film/`, and `/filter-films`.
   - The root route (`/`) serves the main page with the current list of films.
   - The `/add-film/` route processes form submissions to add new films and updates the film list.
   - The `/filter-films` route filters the films based on the user's input and returns the filtered list.

3. **Dynamic Content Rendering**: The Go `html/template` package is used to render the HTML content dynamically, including the film list and individual film elements.

## Installation and Running the Project

To run this project on your local machine, follow these steps:

1. **Clone the Repository**:
   ```bash
   git clone git@github.com:JoelTeoGom/todoLIST-htmx-golang.git
   cd todoLIST-htmx-golang

### Open the Application

- Open your web browser and go to `http://localhost:8000`.
- You should see the film list, filter input, and add film form.

### Interact with the Application

- Use the filter input to search for films by title or director.
- Add new films using the form and see the list update dynamically.

## Dependencies

- **Go (Golang)**: Used for server-side logic and handling HTTP requests.
- **HTMX**: A lightweight JavaScript library used to handle AJAX requests and update parts of the page dynamically.
- **Bootstrap**: A CSS framework used for responsive design and styling.

## License

This project is licensed under the MIT License. Feel free to use and modify it as per your needs.

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or issues, please feel free to open a pull request or issue on the GitHub repository.


   
