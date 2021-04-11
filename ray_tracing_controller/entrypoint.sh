#!/bin/sh
python manage.py migrate
gunicorn --bind 0.0.0.0:8083 --log-leve info --timeout 14400 --workers 100 ray_tracing_controller.wsgi
