FROM python:3.11.7

# Set working directory
WORKDIR /app

RUN apt-get update && apt-get install -y default-mysql-client

# Use requirements for python
COPY ./requirements.txt /app/requirements.txt
RUN pip install --upgrade pip
RUN pip install --no-cache-dir -r requirements.txt

# source code
COPY ./src /app/src

# env
COPY ./.env /app/.env

# alembi.ini file
COPY ./alembic.ini /app/alembic.ini

# entrypoint for migration docker
COPY ./docker/migrate-db.sh /app/docker/migrate-db.sh



# Default command just sleeps (we override it in entrypoint)
CMD ["sleep", "infinity"]
