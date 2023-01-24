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


Call in the terminal in the root of the repo.
```
make run
```

To run tests:
```
make test
```

If you want to launch the app or api only, you can navigate in their respective sub folders and use the same make - commands than in the root directory.


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
