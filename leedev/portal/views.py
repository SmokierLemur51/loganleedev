from django.shortcuts import render

from .models import Project


def projects(request):
    context = {
        "title": "Projects",
        "projects": Project.objects.all(),
    }
    return render(request, "portal/projeccts.html", context)



