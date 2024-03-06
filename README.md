# Oyna Web Application

## Introduction

Welcome to the Oyna Web application! This application is designed to provide a platform for toy rental services, allowing users to explore, rent, and provide feedback on various toys. The codebase is written in Go (Golang) and utilizes the Gorilla web toolkit for routing.

## Table of Contents

- [Home](#home)
- [Feedbacks](#feedbacks)
- [User Management](#user-management)
- [Admin Panel](#admin-panel)
- [Toys](#toys)

## Home <a name="home"></a>

The home endpoint displays the latest toys available for rent. Users can explore the offerings, fostering a dynamic and engaging platform for toy discovery.

## Feedbacks <a name="feedbacks"></a>

- `/create-feedback-form`: Renders a form for users to submit feedback.
- `/create-feedback`: Handles the submission of feedback forms, storing the data in the database.
- `/feedbacks`: Displays a list of all feedbacks submitted by users.
- `/show-feedback/:id`: Shows details of a specific feedback entry.

## User Management <a name="user-management"></a>

- `/signup`: Allows users to sign up for an account, with password validation and email uniqueness checks.
- `/login`: Handles user authentication and redirects authenticated users to the home page.
- `/logout`: Logs out the currently authenticated user.

## Admin Panel <a name="admin-panel"></a>

- `/admin-panel-form`: Renders a form for admin actions, such as user deletion.
- `/admin-panel`: Handles admin actions, such as deleting a user based on their email.

## Toys <a name="toys"></a>

- `/show-toys`: Displays a list of all available toys.
- `/show-toy/:id`: Shows details of a specific toy.
- `/create-toy-form`: Renders a form for users to submit new toys for consideration.
- `/create-toy`: Handles the submission of toy creation forms, storing the data in the database.

## Running the Application

1. Ensure you have Go installed on your machine.
2. Clone the repository: `gh repo clone diorsjoy/toy-rental`
3. Navigate to the project directory: `cd oynaToys`
4. Run the application: `go run main.go`

## Project screenshots
1. Main page / Toys catalogue
![image](https://github.com/diorsjoy/toy-rental/assets/134860120/80782724-d42e-4c13-bc6b-9de54f2d99ae)
2. Toy page
![image](https://github.com/diorsjoy/toy-rental/assets/134860120/451146fc-5b4f-43d5-976e-bf145f3edcee)
3. Feedbacks page
![image](https://github.com/diorsjoy/toy-rental/assets/134860120/c555d7c2-2cdd-427b-bbd3-e0daa1b2daea)
4. Signup page
![image](https://github.com/diorsjoy/toy-rental/assets/134860120/661381d7-772d-41f7-8f21-ab2df1959591)
5. Creating Toy page
![image](https://github.com/diorsjoy/toy-rental/assets/134860120/e494cb41-9c7c-4f4f-af2b-cd04cd0f70c8)
6. Admin page
![image](https://github.com/diorsjoy/toy-rental/assets/134860120/1e3f7ed9-08c9-4651-beb4-04d45ee9d13d)

## Conclusion

Thank you for exploring the Oyna web application. Feel free to contribute, provide feedback, or enhance the features to make it an even more exciting and user-friendly platform!
