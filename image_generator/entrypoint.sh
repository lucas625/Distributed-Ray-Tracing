#!/bin/sh
python manage.py migrate
gunicorn --bind 0.0.0.0:8082 --log-leve info --timeout 14400 --workers 100 image_generator.wsgi
