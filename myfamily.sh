curl -s 'https://learn.zone01kisumu.ke/assets/superhero/all.json' | jq --arg hero "$HERO_ID" ' .[] | select(.id==($hero|tonumber)) |.connections.relatives' | sed 's/"//g'
