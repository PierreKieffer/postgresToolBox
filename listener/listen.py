#!/usr/bin/env python3
import psycopg2
from datetime import datetime
from select import select

def listen() : 
    try : 
        database_config = {
                "deployment" : "custom", 
                "database" : "postgres", 
                "user" : "postgres", 
                "password" : "admin",
                "host" : "172.17.0.2", 
                "port" : "5432"
                }

        connection = connect_database(database_config)
        connection.autocommit = True

        cursor = connection.cursor()
        cursor.execute("LISTEN events;")

        while True : 
            connection.poll()
            while connection.notifies : 
                yield connection.notifies.pop(0)

        cursor.close()

    except (Exception, psycopg2.DatabaseError) as error : 
        print(error)

    finally : 
        if connection is not None : 
            connection.close()

def connect_database(database_config): 

    try : 
        if database_config.get("deployment") == "heroku" : 
            connection = psycopg2.connect(os.environ["DATABASE_URL"], sslmode='require')
            return connection

        else : 

            connection = psycopg2.connect(
            database = database_config.get("database"),
            user = database_config.get("user"),
            password = database_config.get("password"),
            host = database_config.get("host"),
            port = database_config.get("port")
            )

            return connection

    except (Exception, psycopg2.DatabaseError) as error : 
        logger.error(error)

if __name__=="__main__":
    for event in listen() : 
        print(event.payload)

