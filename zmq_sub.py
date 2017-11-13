#!/usr/bin/env python
'''
Created on Jan 5, 2015

@author: gpekht
'''
# Python3 Note: The line below is not needed
#from __future__ import print_function
 
import zmq
import re
import time
 
if __name__ == "__main__":
    context = zmq.Context()
    socket = context.socket(zmq.SUB)
#    socket.setsockopt(zmq.SUBSCRIBE, '#param_p')
#    socket.setsockopt(zmq.SUBSCRIBE, '#config ')
    socket.setsockopt(zmq.SUBSCRIBE, '#pos')
#    socket.setsockopt(zmq.TCP_KEEPALIVE, 1)
    
    # Python3 Note: Use the below line and comment
    # the above line out 
#    socket.setsockopt(zmq.SUBSCRIBE, '')
    #socket.connect("tcp://bo2-ma.rpplabs.com:7001")    
    socket.connect("tcp://192.168.0.100:7001")    
#    socket.connect("tcp://10.7.0.10:7001")
#    socket.connect("tcp://192.168.1.10:7001")
 
    while True:
        msg = socket.recv_multipart()
        # Python3 Note: Use the below line and comment
        # the above line out        
        # msg = socket.recv_string()
#        print("[%s] \n" % (msg))     
#        time.sleep(10)   
        print( msg)
