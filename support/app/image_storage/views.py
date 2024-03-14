from .models import Profile
from .serializers import ProfileSerializer

from rest_framework import permissions, viewsets
from rest_framework.parsers import MultiPartParser, FormParser

class ProfileViewSet(viewsets.ModelViewSet):
    queryset = Profile.objects.all()
    serializer_class = ProfileSerializer
    parser_classes = (MultiPartParser, FormParser)
    permission_classes = [
        permissions.AllowAny
    ]

    def perform_create(self, serializer):
        serializer.save()