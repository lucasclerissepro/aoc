VERSION 0.6

new:
  LOCALLY 

  ARG day
  RUN cp -r .template "$day" 

build:
  LOCALLY 
  FOR dir IN $(ls -d */)
    BUILD "./$dir+build"
  END

run:
  LOCALLY 
  FOR dir IN $(ls -d */)
    BUILD "./$dir+run"
  END

fmt:
  LOCALLY 
  FOR dir IN $(ls -d */)
    BUILD "./$dir+fmt"
  END

clean:
  LOCALLY 
  FOR dir IN $(ls -d */)
    BUILD "./$dir+clean"
  END

bench:
  LOCALLY 
  FOR dir IN $(ls -d */)
    BUILD "./$dir+bench"
  END

docker:
  LOCALLY 
  FOR dir IN $(ls -d */)
    BUILD "./$dir+docker"
  END
