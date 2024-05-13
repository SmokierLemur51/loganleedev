from django.urls import path

from . import views

app_name = "projects"
urlpatterns = [
    path("leads/", views.leads, name="leads"),
    path("leads/<int:pk>", views.lead, name="lead")
    path("clients/", views.clients, name="clients")
    path("clients/<int:pk>", views.client, name="lead")
    path("projects/", views.projects, name="projects"),
    path("projects/<int:pk>/", views.project, name="project"),
]
