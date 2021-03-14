"""Endpoints definition for core app."""

from django.urls import path

from core.views import PathTracingView

urlpatterns = [
    path('path-tracing', PathTracingView.as_view())
]
