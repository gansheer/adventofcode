# from whiplashoo

def apply_mask1(val, mask):
    retval = ''
    for v, m in zip(val, mask):
        if m == '0':
            retval += '0'
        elif m == '1':
            retval += '1'
        else:
            retval += v
    return int(retval, 2)


def apply_mask2(pos, mask):
    x_count = mask.count('X')
    possible_combinations = 2 ** x_count
    ret_arr = [[] for _ in range(possible_combinations)]
    consecutive = 2 ** (x_count - 1)
    for p, m in zip(pos, mask):
        if m == '0':
            for s in ret_arr:
                s.append(p)
        elif m == '1':
            for s in ret_arr:
                s.append('1')
        else:
            wrote = 0
            write_zeros = True
            for s in ret_arr:
                if wrote == consecutive:
                    wrote = 0
                    write_zeros = not write_zeros
                if write_zeros:
                    s.append('0')
                else:
                    s.append('1')
                wrote += 1
            consecutive /= 2
    return [int(''.join(s), 2) for s in ret_arr]


def part_one(inp):
    mem = {}
    for line in inp:
        line = line.strip()
        if 'mask' in line:
            mask = line.split('=')[1][1:]
        else:
            pos = line.split('[')[1].split(']')[0]
            val = line.split('=')[1][1:]
            val = bin(int(val))[2:].zfill(36)
            mem[pos] = apply_mask1(val, mask)
    return (sum(mem.values()))


def part_two(inp):
    mem = {}
    for line in inp:
        line = line.strip()
        if 'mask' in line:
            mask = line.split('=')[1][1:]
        else:
            pos = line.split('[')[1].split(']')[0]
            val = line.split('=')[1][1:]
            pos = bin(int(pos))[2:].zfill(36)
            addresses_to_change = apply_mask2(pos, mask)
            for address in addresses_to_change:
                mem[address] = int(val)
    return (sum(mem.values()))


with open("input_test") as f:
    inp = f.readlines()
    print(part_one(inp))
    print(part_two(inp))
