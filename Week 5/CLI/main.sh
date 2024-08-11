#!/bin/bash

# Create folders
mkdir -p foods drinks snacks

# Create menu.txt files
touch foods/menu.txt drinks/menu.txt snacks/menu.txt

# Add items to foods/menu.txt
echo -e "fried rice\nfried chicken\nchicken porridge" > foods/menu.txt

# Add items to drinks/menu.txt
echo -e "coffee milk\noat milk\nice tea" > drinks/menu.txt

# Fetch snack data from the endpoint and save to snacks/menu.txt
curl -s "https://www.google.com/url?q=https://gist.githubusercontent.com/nadirbslmh/c84ee3527272e52a8e6257d46e627f91/raw/74593cc3fe61ca4ff36e11ed8e3fffcfb1d0c602/snacks.json&sa=D&source=editors&ust=1722834753531745&usg=AOvVaw1aA-qGmfx8xvZInViFscGf" > snacks/menu.txt

