# personality-test-app
Personality test application


## Assignment
Build a simple personality test application, that takes 3-5 different questions, maps them into a score and produces a personality trait of either Introvert or Extrovert.


You can find examples of questions here: https://www.psychologies.co.uk/self/are-you-an-introvert-or-an-extrovert.html


 * An example could be
 * Landing screen
 * Start personality test
 * Dynamic screen, that reads question and answers from a the backend (Just build CRUD, with in memory DB)
 * Finish screen, where you can see your personality trait.


It is a plus if you write some unit-tests.


## How to run this app

Running `./run.sh` should be enough. If you are using a machine, that does not support bash scripts, you can call it's commands separately. For this, you need to have Docker installed on your system.


## Troubleshoot

To check values stored in redis database, open terminal of redis container. You can open redis with command:
```
# You don't need to use --raw, but I recommend in order to pretty print JSON data.
redis-cli --raw
```

You can find json, pretty printed:
```
JSON.GET results INDENT "\t" NEWLINE "\n" SPACE " " 
```