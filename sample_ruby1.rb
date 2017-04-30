#################################################

# 文字列
# "" 特殊文字 式展開
# ''
# puts "hell\no worl\td"
# puts 'hell\no worl\td'

# puts "price #{3000 * 4}"
# puts 'price #{3000 * 4}'

# name = "taguchi"
# puts "hello #{name}"

# + *
puts "hello " + "world"
puts "hello " * 10

#################################################

# - upcase
# - upcase! 破壊的なメソッド
# - downcase reverse

name = "taguchi"
# puts name.upcase
# puts name
# puts name.upcase!
# puts name

# ? 真偽値 true false

p name.empty? # false
p name.include?("g") # true

#################################################

# 配列
colors = ["red", "blue", "yellow"]
# p colors[0] # 添字
# p colors[-1]
# p colors[0..2]
# p colors[0...2]
# p colors[5]  # nil

# colors[0] = "pink"
# colors[1..2] = ["white", "black"]
# colors.push("gold")
# colors << "silver"
# p colors

p colors.size
p colors.sort

#################################################

# ハッシュ
# - key / value

# taguchi 200
# fkoji 400

# scores = {"taguchi" => 200, "fkoji" => 400}
# scores = {:taguchi => 200, :fkoji => 400}
scores = {taguchi: 200, fkoji: 400}

# p scores[:taguchi]
# scores[:fkoji] = 600
# p scores

p scores.size
p scores.keys
p scores.values
p scores.has_key?(:taguchi)
