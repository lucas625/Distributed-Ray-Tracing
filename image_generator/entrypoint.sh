#!/bin/sh
python manage.py migrate
gunicorn --bind 0.0.0.0:8083 --log-leve info --timeout 7200 --workers 4 image_generator.wsgi
