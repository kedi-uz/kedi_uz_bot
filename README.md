## Cheat sheet

To enter a development shell via nix
```shell
nix develop
```


When new position (row) added to the database table it will send webhook or something else to know telegram bot that bot shoudl do some filter action

When new position added telegram bot finds region name and sends notification that registered users belonging only that region

For example: I added region "A" that has a lost pet longitute,latiturde,region name and i registered and set my location as a "A" so in my region has new lost pet so i should get notified about lost pet

# Step by step prosses:

Create ORM init.
Model of USER

User will be registered into db after start and will be asked living District to get notifications
1. username
2. first_name
3. last_name
4. language_code
5. district

## when new lostAnimal notification comes do this
for _ , user := range users {
    if user.distruct == lostAnimal.district then 
    b.SendMessage(int(user.id), lostAnimal details to)
}

# Additional features
0. Create poster for lost pet
1. Get lostAnimal total Count
2. Get lostAnimal total Count by district
3. Get how many users registered (admin only)
4. Broadcast a message
5. Add login via telegram in website like 42.uz otp code comes by phone number
6. Add Points when user subscribers to @kedi_uz channel
7.