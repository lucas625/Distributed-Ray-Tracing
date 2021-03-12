"""Endpoints definition for core app."""

from django.urls import path

from core.views import PngGenerationView

urlpatterns = [
    path('png', PngGenerationView.as_view())
]
