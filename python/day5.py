def int_to_roman(num: int): # missing the 4 and 9 (and 10 multiples of each) proper conversions (i.e. 4 = IV, 9 = IX)
    final = ""
    while num > 0:
        if num / 1000 > 0:
            final += ("M" * (num // 1000))
            num -= ((num // 1000) * 1000)
        if num / 500 > 0:
            final += ("D" * (num // 500))
            num -= ((num // 500) * 500)
        if num / 100 > 0:
            final += ("C" * (num // 100))
            num -= ((num // 100) * 100)
        if num / 50 > 0:
            final += ("L" * (num // 50))
            num -= ((num // 50) * 50)
        if num / 10 > 0:
            final += ("X" * (num // 10))
            num -= ((num // 10) * 10)
        if num / 5 > 0:
            final += ("V" * (num // 5))
            num -= ((num // 5) * 5)
        if num / 1 > 0:
            final += ("I" * (num // 1))
            num -= ((num // 1) * 1)
    
    return final

print(int_to_roman(3))
print(int_to_roman(3999))