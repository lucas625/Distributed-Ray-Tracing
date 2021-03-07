#!/bin/sh
python manage.py migrate
gunicorn --bind 0.0.0.0:8000 --log-leve info --timeout 7200 --workers 4 ray_tracing_controller.wsgi
