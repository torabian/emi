In process of testing, we need the generated code to be imported before the test.
This is egg and chicken issue, but the alternative is create nodevm script which
I couldn't do. We commit these files for visibility in code changes to visually inspect
if generated code makes any sense.