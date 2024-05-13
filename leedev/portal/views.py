from django.shortcuts import get_object_or_404, render

from .models import Lead, Client, Project

def leads(request):
    context = {
        "title": "Leads",
        "leads": Lead.objects.all(),
    }
    return render(request, "portal/leads.html", context)

def lead(request, lead_id):
    lead = get_object_or_404(Lead, pk=lead_id)
    context = {
        "title": lead.company,
        "lead": lead,
    }
    return render(request, "portal/lead.html", context)

def clients(request):
    context = {
        "title": "Clients",
        "clients": Lead.objects.all(),
    }
    return render(request, "portal/clients.html", context)

def client(request, client_id):
    client = get_object_or_404(Client, pk=client_id)
    context = {
        "title": client.company,
        "client": client,
    }
    return render(request, "portal/client.html", context)


def projects(request):
    context = {
        "title": "Projects",
        "projects": Project.objects.all(),
    }
    return render(request, "portal/projects.html", context)

def project(request, project_id):
    context = {
        "title": "Project",
        "projects": get_object_or_404(Project, pk=project_id),
    }
    return render(request, "portal/project.html", context)


