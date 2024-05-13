from django.shortcuts import render, get_object_or_404

def index(request):
    # projects & portfolio info
    context = {"title": "Public"}
    return render(request, "public/index.html", context)


def portfolio(request):
    context = {"title": "Portfolio"}
    return render(request, "public/portfolio.html", context)


def services(request):
    context = {"title": "Services"}
    return render(request, "public/services.html", context)


def new_contact_request(request):
    pass

