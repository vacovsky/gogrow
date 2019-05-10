# gogrow

for monitoring an enclosed plant-growing environment

## screenshot
![Alt text](/static/images/screenshot.png "Screenshot")

## setup

```bash
go get github.com/vacovsky/gogrow/
go get github.com/vacovsky/gorm

cd static
npm install
cd ..
# if you want the webcam stuff to function:
apt-get install fswebcam 

# edit the launch file with env variables as needed
sudo ./launch.sh
```

## features

- takes webcam image at interval
- tracks all kind of useful data, and charts a few
- very easy to set up the hardware (uses 2 DHT11 senors plugged directly into the Pi3)
- will add soil thermometer stuff later

## notes

- probably going to need to run with sudo `sudo launch.sh`
- works great for peppers
- took a many shortcuts to be able to crank this out in a couple evenings