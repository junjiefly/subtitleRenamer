# subtitleRenamer
eg:
	rename for a tv series
		
[root@host The Killing Vote]# ls -a

-rwxrwxrwx 1 nginx nginx 42345 8 月 12 00:26 backdrop1.jpg

-rwxrwxrwx 1 nginx nginx 42315 8 月 12 00:26 folder.jpg

drwxrwxrwx 2 nginx nginx 77 8 月 11 20:43 metadata

-rwxrwxrwx 1 nginx nginx 78049 8 月 12 00:26 **The.Killing.Vote.S01E01.1080p.AMZN.WEB-DLA.CSD.ES.0.H.265.cht.srt**

-rwxrwxrwx 1 nginx nginx 1165027229 8 月 11 18:24 *The.Killing.Vote.S01E01.1080p.AMZN.WEB-DL.DDP2.0.H.265-thon.mkv*

-rwxrwxrwx 1 nginx nginx 3830 8 月 12 00:26 The.Killing.Vote.S01E01.1080p.AMZN.WEB-DL.DDP2.0.H.265-thon.nfo




do rename:

[root@host The Killing Vote]# ./subtitleRenamer 

find episode:   1, subtitle matched? yes 



changes:
  [root@host The Killing Vote]# ls -a
  

-rwxrwxrwx 1 nginx nginx      42345 8月  12 00:26 backdrop1.jpg

-rwxrwxrwx 1 nginx nginx      42315 8月  12 00:26 folder.jpg

drwxrwxrwx 2 nginx nginx         77 8月  11 20:43 metadata

-rwxrwxrwx 1 nginx nginx      78049 8月  12 00:26 **The.Killing.Vote.S01E01.1080p.AMZN.WEB-DL.DDP2.0.H.265-thon.cht.srt**

-rwxrwxrwx 1 nginx nginx 1165027229 8月  11 18:24 *The.Killing.Vote.S01E01.1080p.AMZN.WEB-DL.DDP2.0.H.265-thon.mkv*

-rwxrwxrwx 1 nginx nginx       3830 8月  12 00:26 The.Killing.Vote.S01E01.1080p.AMZN.WEB-DL.DDP2.0.H.265-thon.nfo

-rwxrwxrwx 1 nginx nginx       3046 8月  12 00:26 tvshow.nfo



