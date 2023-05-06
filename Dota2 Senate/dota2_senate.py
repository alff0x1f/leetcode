class Solution:
    def find_next(self, senate: list, start_position: int, party: str) -> int:
        pos = start_position
        while True:
            pos += 1
            if pos >= len(senate):
                pos = 0
            if senate[pos] == party:
                return pos
            if pos == start_position:
                return -1

    def predictPartyVictory(self, senate: str) -> str:
        senate = [senator for senator in senate]
        radiant = -1
        dire = -1

        # find position
        for i, senator in enumerate(senate):
            if senator == "R" and radiant < 0:
                radiant = i
            elif senator == "D" and dire < 0:
                dire = i
            if radiant > 0 and dire > 0:
                break

        # simple case, only one party
        if radiant < 0:
            return "Dire"
        if dire < 0:
            return "Radiant"

        current = 0
        while True:
            print(''.join(senate), current, senate[current], radiant, dire)
            if senate[current] == "R":
                if radiant == current:
                    radiant = self.find_next(senate, radiant, "R")
                # vote
                senate[dire] = "-"
                dire = self.find_next(senate, dire, "D")
                if dire < 0:
                    return "Radiant"
            elif senate[current] == "D":
                if dire == current:
                    dire = self.find_next(senate, dire, "D")
                # vote
                senate[radiant] = "-"
                radiant = self.find_next(senate, radiant, "R")
                if radiant < 0:
                    return "Dire"
            current += 1
            if current >= len(senate):
                current = 0


if __name__ == "__main__":
    s = Solution()
    print(s.predictPartyVictory("RDD"))


