#!/bin/sh

Message=$2

generate_post_data()
{
  cat <<EOF
    {   
        "message":"$Message",
        "title":"Zabbix Notification"
    }
EOF
}

curl --location --request POST "localhost:9999/send/notif?to=$1" \
--header 'Content-Type: text/plain' \
--data-raw "$(generate_post_data)"


echo "$(date '+%Y-%m-%d %H:%M:%S')  SMS Text sent to $1 Message:$Message" >> sms-send.log
