# subtitleRenamer
eg1:
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





eg2:
	rename for a movie
		
[root@host ironMan1]# ls -a

-rwxrwxrwx  1 nginx nginx     102566 7月  24 23:48 folder.jpg

-rwxrwxrwx. 1 nginx nginx 1773975793 7月  22 13:11 Iron.Man.2008.1080p.BrRip.x264.YIFY.mp4

-rwxrwxrwx  1 nginx nginx       6275 8月   9 15:03 Iron.Man.2008.1080p.BrRip.x264.YIFY.nfo

-rwxrwxrwx  1 nginx nginx     270900 7月  31 21:59 **Iron.Man.2008.zh.ass**




do rename:

[root@host ironMan1]# ./subtitleRenamer 

find move: Iron.Man.2008.1080p.BrRip.x264.YIFY, subtitle matched? yes


changes:

  [root@host The Killing Vote]# ls -a

-rwxrwxrwx  1 nginx nginx     102566 7月  24 23:48 folder.jpg

-rwxrwxrwx. 1 nginx nginx 1773975793 7月  22 13:11 Iron.Man.2008.1080p.BrRip.x264.YIFY.mp4

-rwxrwxrwx  1 nginx nginx       6275 8月   9 15:03 Iron.Man.2008.1080p.BrRip.x264.YIFY.nfo

-rwxrwxrwx  1 nginx nginx     270900 7月  31 21:59 **Iron.Man.2008.1080p.BrRip.x264.YIFY.zh.ass**

-rwxrwxrwx  1 nginx nginx     123419 7月  24 23:48 logo.png





