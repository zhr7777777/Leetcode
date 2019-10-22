class Solution {
    public String reverseVowels(String s) {
        Set<Character> vowels = new HashSet<Character>();
        vowels.add('a');
        vowels.add('e');
        vowels.add('i');
        vowels.add('o');
        vowels.add('u');
        vowels.add('A');
        vowels.add('E');
        vowels.add('I');
        vowels.add('O');
        vowels.add('U');
        List<Character> vowelArr = new ArrayList<Character>();
        char[] sArr = s.toCharArray();
        for(int i=0; i<s.length(); i++) {
            char cur = s.charAt(i);
            if(vowels.contains(cur)) {
                vowelArr.add(cur);
            }
        }
        int index = vowelArr.size() - 1;
        for(int i=0; i<sArr.length; i++) {
            if(vowels.contains(sArr[i])) {
                sArr[i] = vowelArr.get(index);
                index--;
            }
        }
        return new String(sArr);
    }
}