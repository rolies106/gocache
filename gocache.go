package gocache

import (
	"fmt"
)

func SetMultiKeysCache(cache_type string, keys map[string]interface{}, value string, no_expire ...bool) {
	fmt.Println("%+v", "ConnectionSuccess")
	fmt.Println("%+v", Connection)
}

func GetCacheKey(cache_type string, keys map[string]interface{}) string {

	key := ""
	
	switch cache_type {
	case "value":
		// Special case
	default:
		if len(map) == 0 {
			key = cache_type
		} else {

		}
	}

            if (!empty($parameters)) {

                $key = $key . $type;

                foreach ($parameters as $k => $value)
                {
                    $key = $key . '_' . $k . '_' . $value;
                }

            } else {
                $key = $type;
            }

    }

    return md5($key);
}


        $expire = [];
        $cache_key = $this->getCacheKey($type, $keys);

        if ($no_expire == false)
            $expire = ['nx', 'ex' => $this->ttl];

        // Cache the value
        $this->instance->set($cache_key, $value, $expire);

        // Assign this cache key as index members
        $this->setIndexMembers($type, $keys, $cache_key);