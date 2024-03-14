from rest_framework import serializers
from .models import Profile

class ProfileSerializer(serializers.ModelSerializer):

    avatar = serializers.ImageField(required=False)

    class Meta:
        model = Profile
        fields = ['avatar']