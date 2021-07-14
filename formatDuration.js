function formatDuration(seconds) {
    const dict = {
        year : 31536000,
        day: 86400,
        hour: 3600,
        minute: 60,
        second: 1
    }

    let result = ''

    Object.keys(dict).forEach(key => {
        const number = Math.floor(seconds / dict[key])
        seconds = seconds % dict[key]
        const string = number > 0 ? (number + ' ' + key + (number > 1 ? 's' : '')) : ''
        result += string + (number > 0 && seconds > 0 ? ', ' : '')
    })

    return result.replace(/,(?=[^,]*$)/, ' and') || 'now'
}

console.log(formatDuration(3722))