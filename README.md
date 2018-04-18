
# Goals
- generate SSH config files for servers spread across multiple enviornments and stacks (configure)
	- ensure that config files support multiple users
	- ensure that config files support multiple environments
	- enusre that config files support multiple stacks per environment
	- allow for different SSH keys for each stack/environment.
	- allow for different jumphost/bastion domain generation template for each env.
	- handle multiple bastion hosts per stack (eg. range 1-X)

- allow for a method to retrieve config file with stack/env/user defaults (list)
- delete SSH config files & other data from host (cleanup)
- the ability to download files from hosts, though the tunnel (scp)
- the ability to open the ssh tunnel, with http port binding locally (connect)
	- local ports chosen will be dynamic and depend on the hash of the config filepath (unique on the config level) https://stats.stackexchange.com/questions/26344/how-to-uniformly-project-a-hash-to-a-fixed-number-of-buckets
	- the ability to create/update a pac file, which points to a proxy server inside behind the bastion (--pac)

- ability to hook into the configure/create step and create additional templates dynamically on the file system.
	- used to create knife configs
	- used to create proxy files?



# How does config file work:

- list of questions, each question has a key, that key is used to populate the template.
- each question has a description, used when asking the user for input
- each question has a type, a golang type, used when generating the struct, and for validaiton
- question can have validation, ensure that it's proper value
	- https://github.com/go-playground/validator
	- https://github.com/bluesuncorp/validator
	- https://github.com/xeipuuv/gojsonschema
	- https://github.com/thedevsaddam/govalidator
	- https://github.com/go-validator/validator
	- https://github.com/gima/govalid
	- https://github.com/lestrrat/go-jsref
	- https://medium.com/@lestrrat/json-schema-and-go-3c7439959077
	- https://github.com/lestrrat/go-jsschema

- question can have range of allowed values
- question can have an example string (not default), used for hinting to the user.
- question can have ui_group_by value, 1,2,3 used in ui for listing.
- question can have ui_hidden value, boolean, used in ui to hide during listing.

- questions will be used to create a dynamic Struct, with tags added dynamically: https://github.com/fatih/gomodifytags

- list of answers
- answers can reference an external file using `_file`, which will be loaded inplace.
- answers must provide atleast one of the questions. (empty objects will throw an error)
- answers will be validated against the questions. Any invalid answers removed? throw an error?

- template section
- custom/overridable templates supported:
	- config template
	- config filename template
	- ssh key filepath template


# References

- https://github.com/moul/awesome-ssh/blob/master/README.md
- https://github.com/dbrady/ssh-config
- https://github.com/k4m4/terminals-are-sexy
- https://github.com/n1trux/awesome-sysadmin
- https://github.com/cjbarber/ToolsOfTheTrade
- https://github.com/dastergon/awesome-sre
- https://stackoverflow.com/questions/17355667/replace-current-process
- https://stats.stackexchange.com/questions/26344/how-to-uniformly-project-a-hash-to-a-fixed-number-of-buckets
- moul/advanced-ssh-config
- https://github.com/emre/storm
- https://stackoverflow.com/questions/12484398/global-template-data
- https://stackoverflow.com/questions/35612456/how-to-use-golang-template-missingkey-option
- https://medium.com/@dgryski/consistent-hashing-algorithmic-tradeoffs-ef6b8e2fcae8
- https://github.com/mitchellh/go-homedir