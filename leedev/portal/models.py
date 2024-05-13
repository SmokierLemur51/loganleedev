from django.db import models
from django.utils import timezone


class ProjectStatus(models.Model):
    status = models.CharField(max_length=75)
    description = models.TextField()
    
    def __str__(self):
        return self.status


class Lead(models.Model):
    company = models.CharField(max_length=100)
    point_of_contact = models.CharField(max_length=60, null=True)
    email = models.EmailField(null=True)
    phone = models.CharField(max_length=15)
    website = models.URLField(null=True) # use the filter by to search for facebook 
    converted = models.BooleanField(default=False)

    def __str__(self):
        return self.company

    def convert_to_client(self):
        if not self.converted:
            c = Client(
                company=self.company, point_of_contact=self.point_of_contact, 
                email=self.email, phone=self.phone, website=self.website
            )
            c.save()
            self.converted = True
            self.save()
            return c
        else:            
            raise ValueError("Lead has already been converted to client")


class Client(models.Model):
    company = models.CharField(max_length=100)
    point_of_contact = models.CharField(max_length=60, null=True)
    email = models.EmailField(null=True)
    phone = models.CharField(max_length=15, null=True)
    website = models.URLField(null=True)

    def __str__(self):
        return self.company


class ClientNote(models.Model):
    client = models.ForeignKey(Client, on_delete=models.CASCADE)
    title = models.CharField(max_length=100)
    content = models.TextField(null=True)
    created_at = models.DateTimeField(default=timezone.now)
    updated_at = models.DateTimeField(auto_now=True)


class Project(models.Model):
    title = models.CharField(max_length=100)
    description = models.TextField()
    total_time_spent = models.FloatField(default=0.0)
    hourly_rate = models.FloatField(default=0.0)
    wholesale = models.FloatField(default=0.0)
    subscription = models.FloatField(default=0.0)
    client = models.ForeignKey(Client, on_delete=models.CASCADE)
    status = models.ForeignKey(ProjectStatus, on_delete=models.CASCADE)
    created_at = models.DateTimeField(default=timezone.now)
    updated_at = models.DateTimeField(auto_now=True)

    def __str__(self):
        return f"Project: {self.title}, Company: {self.client.company}"


class TicketStatus(models.Model):
    status = models.CharField(max_length=75)
    description = models.TextField()

    def __str__(self):
        return self.status


class ProjectTicket(models.Model):
    project = models.ForeignKey(Project, on_delete=models.CASCADE)
    status = models.ForeignKey(TicketStatus, on_delete=models.CASCADE)
    title = models.CharField(max_length=100)
    description = models.TextField(null=True)
    created_at = models.DateTimeField(default=timezone.now)
    updated_at = models.DateTimeField(auto_now=True)

    def __str__(self):
        return f"Project: {self.project.title} Ticket: {self.title}"


class ProjectNote(models.Model):
    project = models.ForeignKey(Project, on_delete=models.CASCADE)
    title = models.CharField(max_length=100)
    content = models.TextField(null=True)
    created_at = models.DateTimeField(default=timezone.now)
    updated_at = models.DateTimeField(auto_now=True)


class ProposalProject(models.Model):
    # these will be used to create endpoints for potential customers to see
    title = models.CharField(max_length=100)
    description = models.TextField()
    total_time_spent = models.FloatField(default=0.0)
    hourly_rate = models.FloatField(default=0.0)
    wholesale = models.FloatField(default=0.0)
    subscription = models.FloatField(default=0.0)
    lead = models.ForeignKey(Lead, on_delete=models.CASCADE)
    status = models.ForeignKey(ProjectStatus, on_delete=models.CASCADE)
    created_at = models.DateTimeField(default=timezone.now)
    updated_at = models.DateTimeField(auto_now=True)