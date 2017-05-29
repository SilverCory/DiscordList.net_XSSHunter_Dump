# XSSDump

### Why?
2 months have now passed and it's not fixed. Not only that but the XSSHunter script that's on there has not been removed! It's causing a large amount of spam to my email.

### What?
XSS (Cross Site Scripting) is a exploit caused by unsanitised user inputs. It allows for html to be injected on a page, that's fine and not really insecure for normal HTML tags, however it also allows for Javascipt to be ran.

Javascript has the ability to do so much, and pretend to be you. From [deleting all the bots](https://www.youtube.com/watch?v=QvXxtSNi8Jc&t=50), to a simple alert. Not only that, but privacy is compromised.

### How to fix it?
For different languages, there's different ways. I, and a few others, have discussed ways of fixing this with the developer/owner of this site. This was 2 months ago.

**ALL USER INPUT SHOULD BE SANATISED AND CHECKED TO ENSURE THAT THE INPUT IS WHAT IS EXPECTED AND ALLOWED**. Even predefined dropdown's aren't safe!

### What can I do?
- If you use this site?
  - URGE THE OWNER TO FIX THIS ISSUE!

- If you fear your site may be vulnerable?
  - Read up on the latest articles on website security and XSS.
  - Get your site tested.
  - Take backups (as you should anyway!)

I'm no security professional, I'm a developer, so don't take my advice seek someone who knows a lot on this subject if you really want details.

There's also a few headers, and cookie flags you can use, like `same origin header`, and `secure cookie flag`. These will help minimise compromised data.

### Who?
- author of the site:
  - I won't shame them yet. Another month passes and I will sure.
- me:
  - some noob..
  - Literally a child could do these types of attacks.


### Where?
http://discordlist.net/

### When?
... Commit history?

### How?
We're blessed by XSSHunter for their awesome app. Check it out! You even get payloads to bypass some of the most basic XSS prevention filters.
