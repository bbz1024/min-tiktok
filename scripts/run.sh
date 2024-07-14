#!/bin/ash

function clear_and_run() {
    local name=$1
    local type=$2

    # Capitalize the first letter of the name
    local capName="${name^}"

    # Check if the type is 'api'
    if [ "$type" == "api" ]; then
        local capNameApi="${capName}Api"
        echo "clearing api"
        pushd api || exit
        # Loop through all directories and remove those not matching the name
        for file in *; do
            if [ "$file" != "$capNameApi" ]; then
                echo "clearing $file"
                rm -rf "$file"
            fi
        done
        popd || exit
        # Run the specified api
        echo "running $capNameApi"
        ./api/"$capNameApi"
    fi

    # Check if the type is 'service'
    if [ "$type" == "service" ]; then
       local capNameService="${capName}Service"
        echo "clearing service"
        pushd services || exit
        # Loop through all directories and remove those not matching the name
        for file in *; do
            if [ "$file" != "$capNameService" ]; then
                echo "clearing $file"
                rm -rf "$file"
            fi
        done
        popd || exit
        # Run the specified service
        echo "running $capNameService"
        ./services/"$capNameService"
    fi
}


clear_and_run "$1" "$2"
