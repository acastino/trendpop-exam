# Introduction

This is my implementation of the assessment task that Trendpop gave me.

Please visit each of the folders, to find specific instructions on that part of the system.

Thank you so much for taking the time.

Kind regards.

# Usage

1. The frontend is currently deployed to: https://acastino.github.io/trendpop-exam/
   - You can switch the Backend API to Golang, by adding the following query string: `?backend=golang`
   - The default Backend is Laravel, which optionally can use the following query string: `?backend=laravel`
1. The backend servers are currently deployed to:
   1. Laravel: https://trendpop-exam-laravel-production.up.railway.app/
   1. Golang: https://trendpop-exam-golang-production.up.railway.app/

## Notes

- I have added an endpoint that lists the latest 10 submits, in descending order, both for Laravel and Golang, at the following path: `/api/invoice`. You should be able to verify that your data has been saved properly, using these endpoints.

- Both backend services use the same database behind the scenes. So, if you add an entry on the frontend using Golang as your backend, then you should be able to use the "list" endpoint on the Laravel side of things (in addition to Golang's "list" endpoint), and expect to see your latest entry there, as well; and vice versa.
