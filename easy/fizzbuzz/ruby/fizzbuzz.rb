def fizzbuzz(firstDivisor, secondDivisor, n)
	n.times do |currentInteger|
		currentInteger = currentInteger + 1
		output = ""
		firstRemainder = currentInteger % firstDivisor
		secondRemainder = currentInteger % secondDivisor

		if firstRemainder == 0
			output += "F"
		end

		if secondRemainder == 0
			output += "B"
		end

		if output == ""
			output = currentInteger.to_s
		end

		print output + " "
	end

	puts
end

File.open(ARGV[0]).each_line do |line|
	values = line.split(" ")
	fizzbuzz values[0].to_i, values[1].to_i, values[2].to_i
end