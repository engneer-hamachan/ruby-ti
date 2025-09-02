# def test
#   a = {}
#   a[:b] = 1
# 
#   a
# end
# 
# test[:b] = '11111'
# 
# # isMethodで良い気がする。あとtest148が地味におかしい
# # 式終わりまで評価する必要がある
# # test148はIsBindの分岐が、a = h.test[0] にたいおうできてない
# dbtp test[:b]


# Array版
def test
  a = []
  a[0] = 1

  a
end


test[0] = '11111'
dbtp test[0]
