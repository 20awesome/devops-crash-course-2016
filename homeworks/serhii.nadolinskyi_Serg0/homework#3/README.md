1. post how you blocked acces to folder:
   ```
    chmod 000 vlad_dir/
    ```

2. file outside home dir : it's name ,where it's placed, what purpose of dir ( where this file is located):
    ```
    sudo find / -type f -newermt 2016-08-02 | grep -Ev '/dev|/proc|/tmp|/run'
    ```
    files: a lot of and __/etc/touched__ . I think it is =)
    it located in __/etc__ folder, it contains Host-specific system-wide configuration files

3. [my image]: https://52.28.184.172/dev-ops-bucket/docker.png

