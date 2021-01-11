import os

BASE_DIR = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

MICRO_DIR = BASE_DIR + '/micro'

SERVICES_DIR = MICRO_DIR + '/service'

POCSUITE_SRV_DIR = SERVICES_DIR + '/poc_scanner'

POC_PLUGINS_DIR = POCSUITE_SRV_DIR + '/poc_plugins'
