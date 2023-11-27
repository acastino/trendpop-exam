<p align="center"><a href="https://laravel.com" target="_blank"><img src="https://raw.githubusercontent.com/laravel/art/master/logo-lockup/5%20SVG/2%20CMYK/1%20Full%20Color/laravel-logolockup-cmyk-red.svg" width="400" alt="Laravel Logo"></a></p>

<p align="center">
<a href="https://github.com/laravel/framework/actions"><img src="https://github.com/laravel/framework/workflows/tests/badge.svg" alt="Build Status"></a>
<a href="https://packagist.org/packages/laravel/framework"><img src="https://img.shields.io/packagist/dt/laravel/framework" alt="Total Downloads"></a>
<a href="https://packagist.org/packages/laravel/framework"><img src="https://img.shields.io/packagist/v/laravel/framework" alt="Latest Stable Version"></a>
<a href="https://packagist.org/packages/laravel/framework"><img src="https://img.shields.io/packagist/l/laravel/framework" alt="License"></a>
</p>

## Introduction

This backend service is using Laravel, so that we can make use of Eloquent and the Database Migration features of Laravel, among other shortcuts and automations, mostly due to time constraints.

For a more specific use-case (maybe when not using MySQL), then perhaps another programming language altogether would be a better choice.

Or if you want a much better performance, then maybe a compiled language like Golang, Java, Rust, etc would be the much better option for that kind of requirement, although the tooling would all be separated unlike what we have in Laravel.

## Project setup (development)

Edit the `.env` after copying it from example.

```
cp .env.example .env
composer install
php artisan config:clear
php artisan config:cache
php artisan key:generate
php artisan migrate
php artisan serve
```

## Project setup (production)

A lot of tools can automate the deployment of Laravel, so prefer to use those, by default. Remember to perform the following however, to make sure that your app is production ready.

1. Edit the `.env` file, specifically the variables that start with `APP_`.
1. Generate a unique key for this new instance of your app, using `php artisan key:generate`.
1. Prepare the database tables, by running: `php artisan migrate`

## Security Vulnerabilities

If you discover a security vulnerability within Laravel, please send an e-mail to Taylor Otwell via [taylor@laravel.com](mailto:taylor@laravel.com). All security vulnerabilities will be promptly addressed.

## License

The Laravel framework is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).
