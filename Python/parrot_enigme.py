#!/usr/bin/python

import base64
import binascii

# http://www.commentcamarche.net/contents/94-codage-base64
# explication du signe = à la fin...
enigme="MjV4MjUg/ka/wW6QbpwLt1r126Kq7BU9B/qq/gAWAPvv1RYhCKi2V3uOex\
d9Gr/PZKpbn3camSmmKvoAUMc/vivwRnGbqx/91Hte6ppLBdv5/upfgA=="
print(binascii.a2b_base64(enigme))
#print("cGFycm90" in enigme)
#hex_test=base64.b64decode(enigme)
#hex_text_bis=binascii.hexlify(hex_test)
#hex_text_bis.strip()
#print(binascii.a2b_hex(hex_text_bis))
