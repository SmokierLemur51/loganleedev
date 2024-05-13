from django.db import models

class ContactRequest(models.Model):
    name = models.CharField(max_length=200)
    phone = models.CharField(max_length=20)
    email = models.EmailField(null=True) # can be null, some people dont have email
    information = models.TextField()

